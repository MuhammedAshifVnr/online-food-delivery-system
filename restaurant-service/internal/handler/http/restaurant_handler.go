package http

import (
	"net/http"
	"strconv"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/app/dto"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/repository"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/pkg/database"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/pkg/response"
	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	service *service.RestaurantService
}

func NewRestaurantHandler(service *service.RestaurantService) *RestaurantHandler {
	return &RestaurantHandler{service}
}

func (h *RestaurantHandler) AddRestaurant(c *gin.Context) {
	var restaurantDTO dto.RestaurantDTO
	if err := c.BindJSON(&restaurantDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant, err := h.service.CreateRestaurant(restaurantDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": restaurant})
}

func (h *RestaurantHandler) GetRestaurant(c *gin.Context) {
	restaurant, err := h.service.GetRestaurantByID()
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, restaurant, "")
}

func (h *RestaurantHandler) AddMenuItem(c *gin.Context) {
	var menuItemDTO dto.MenuItemDTO
	if err := c.BindJSON(&menuItemDTO); err != nil {
		response.JSON(c, http.StatusBadRequest, nil, err.Error())
		return
	}
	resId, _ := strconv.Atoi(c.Param("id"))
	menuItemDTO.RestaurantID = uint(resId)

	menuItem, err := h.service.CreateMenuItem(menuItemDTO)
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.JSON(c, http.StatusCreated, menuItem, "")
}

func (h *RestaurantHandler) GetMenuItems(c *gin.Context) {
	restaurantID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, nil, "Invalid restaurant ID")
		return
	}

	menuItems, err := h.service.GetMenuItemsByRestaurantID(uint(restaurantID))
	if err != nil {
		response.JSON(c, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, menuItems, "")
}

func RegisterRestaurantRoutes(router *gin.Engine) {
	restaurantHandler := NewRestaurantHandler(service.NewRestaurantService(repository.NewRestaurantRepository(database.GetDB())))

	restaurantRoutes := router.Group("/restaurants")
	{
		restaurantRoutes.POST("/", restaurantHandler.AddRestaurant)
		restaurantRoutes.GET("/", restaurantHandler.GetRestaurant)
		restaurantRoutes.POST("/:id/menu-items", restaurantHandler.AddMenuItem)
		restaurantRoutes.GET("/:id/menu-items", restaurantHandler.GetMenuItems)
	}
}
