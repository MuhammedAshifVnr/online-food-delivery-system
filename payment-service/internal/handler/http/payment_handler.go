package http

import (
	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: service}
}

func (h *PaymentHandler) RegisterRoutes(router *gin.Engine) {
	PaymentGroup := router.Group("/payment")

	PaymentGroup.POST("/confirm",h.PaymentConfirm)
	PaymentGroup.GET("/complete",h.PaymentComplete)
}

func(h *PaymentHandler) PaymentConfirm(c *gin.Context){
	h.paymentService.PaymentConfirmation(c)
}

func (h *PaymentHandler) PaymentComplete(c *gin.Context){
	h.paymentService.PaymentComplete(c)
}