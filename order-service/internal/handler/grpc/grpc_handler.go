package grpc

import (
	"context"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/proto"
)

type GrpcHnadler struct {
	service service.OrderService
	proto.UnimplementedOrderServiceServer
}

func NewOrderGrpcHnadler (service service.OrderService)GrpcHnadler{
	return GrpcHnadler{service: service}
}

func (g *GrpcHnadler) CompletePayment(ctx context.Context, req *proto.PaymentCompleteRequest)(*proto.PaymentCompleteResponse,error){
	status, err :=g.service.CompletePayment(req.OrderId)
	return &proto.PaymentCompleteResponse{
		Status: status,
	},err
}
