package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/app/service"
	pb "github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/proto"
)

type PaymentHandler struct {
	service service.PaymentService
	pb.UnimplementedPaymentServiceServer
}

func NewGrpcPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) NewOrder(ctx context.Context, req *pb.NewOrderRequest) (*pb.NewOrderResponse, error) {
	fmt.Println("entered")
	razorOrderId, err := h.service.NewOrder(req.OrderId, req.Price)
	if err != nil {
		log.Fatal(err)
	}
	return &pb.NewOrderResponse{
		RazorOrderId: razorOrderId,
	}, nil
}
