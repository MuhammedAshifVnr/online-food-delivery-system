package main

import (
	"log"
	"net"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/app/service"
	grpcHnadler "github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/handler/grpc"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/handler/http"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/internal/repository"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/pkg/database"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/payment-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	database.Init()

	orderConn, err := grpc.Dial("order-service:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to restaurant service: %v", err)
	}
	defer orderConn.Close()
	orderClient := proto.NewOrderServiceClient(orderConn)

	repo := repository.NewPaymentRepository(database.GetDB())
	service := service.NewPaymentService(repo, orderClient)
	handler := http.NewPaymentHandler(service)

	router := gin.Default()
	router.LoadHTMLGlob("/root/pkg/temp/*.html")
	handler.RegisterRoutes(router)

	go func() {
		if err := router.Run(":8083"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	GrpcHandler := grpcHnadler.NewGrpcPaymentHandler(service)
	grpcServer := grpc.NewServer()
	proto.RegisterPaymentServiceServer(grpcServer, GrpcHandler)

	lis, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatalf("Failed to listen on port 9093: %v", err)
	}

	log.Println("gRPC server is running on port 9093")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
