package http

import (
	"net/http"
	"strconv"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/app/dto"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func RegisterOrderRoutes(router *gin.Engine, orderService *service.OrderService) {
	handler := NewOrderHandler(orderService)
	orderGroup := router.Group("/orders")
	orderGroup.Use(middleware.AuthMiddleware())
	{
		orderGroup.POST("", handler.PlaceOrder)
		orderGroup.PUT("/:id/status", handler.UpdateOrderStatus)
		orderGroup.GET("/:id", handler.GetOrderDetails)
		orderGroup.GET("/history", handler.GetOrderHistory)
		orderGroup.GET("", handler.GetAllOrders)
	}
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	req.UserID = c.GetUint("userID")

	createdOrder, err := h.orderService.PlaceOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error2": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var req dto.UpdateOrderStatusRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.orderService.UpdateOrderStatus(uint(orderID), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Order status updated"})
}

func (h *OrderHandler) GetOrderDetails(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.orderService.GetOrderDetails(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrderHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	orders, err := h.orderService.GetOrderHistory(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
