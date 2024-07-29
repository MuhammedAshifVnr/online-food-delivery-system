package main

import (
	"log"
	"net"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/handler/http"
	grpcHnadler "github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/handler/grpc"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/repository"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/pkg/database"
	proto "github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	database.Init()

	userConn, err := grpc.Dial("user-service:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()
	userClient := proto.NewUserServiceClient(userConn)

	restaurantConn, err := grpc.Dial("restaurant-service:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to restaurant service: %v", err)
	}
	defer restaurantConn.Close()
	restClient := proto.NewRestaurantServiceClient(restaurantConn)

	payConn, err := grpc.Dial("payment-service:9093", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to restaurant service: %v", err)
	}
	defer restaurantConn.Close()
	payClient := proto.NewPaymentServiceClient(payConn)

	orderRepo := repository.NewOrderRepository(database.GetDB())
	orderService := service.NewOrderService(orderRepo, userClient, restClient, payClient)

	router := gin.Default()
	http.RegisterOrderRoutes(router, orderService)

	go func() {
		if err := router.Run(":8082"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()


	GrpcHandler := grpcHnadler.NewOrderGrpcHnadler(*orderService)
	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer,&GrpcHandler)
	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Fatalf("Failed to listen on port 9092: %v", err)
	}

	log.Println("gRPC server is running on port 9092")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
