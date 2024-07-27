package service

import (
	"context"
	"log"
	"net/http"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/repository"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/proto"
	"github.com/gin-gonic/gin"
)

type PaymentService interface {
	PaymentConfirmation(c *gin.Context)
	PaymentComplete(c *gin.Context)
	NewOrder(orderId string, price uint32) (string, error)
}

type paymentService struct {
	repo        repository.PaymentRepository
	orderClient proto.OrderServiceClient
}

func NewPaymentService(repo repository.PaymentRepository,order proto.OrderServiceClient) *paymentService {
	return &paymentService{repo: repo,orderClient: order}
}

func (s *paymentService) PaymentConfirmation(c *gin.Context) {
	orderID, status := s.repo.PaymentConfirmation(c)
	if status != "success" {
		log.Fatal("Payment failed")
	}

	if s.orderClient == nil {
		log.Println("orderClient is not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	
	_, err := s.orderClient.CompletePayment(context.Background(), &proto.PaymentCompleteRequest{
		OrderId: orderID,
	})
	if err != nil {
		log.Fatal("failed to update booking data")
	}

}

func (p *paymentService) PaymentComplete(c *gin.Context) {
	c.HTML(200, "payment.html", nil)
}

func (p *paymentService) NewOrder(orderId string, price uint32) (string, error) {
	return p.repo.NewRazorOrder(orderId, price)
}
