package main

import (
	"log"
	"net"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/handler/http"
	grpcHandler "github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/handler/grpc"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/repository"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/pkg/database"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/pkg/middleware"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	database.Init()

	router := gin.Default()

	router.Use(middleware.AuthMiddleware())
	http.RegisterRestaurantRoutes(router)

	
	go func() {
		if err := router.Run(":8081"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	grpcServer := grpc.NewServer()

	restaurantRepository := repository.NewRestaurantRepository(database.GetDB())
	restaurantService := service.NewRestaurantService(restaurantRepository)
	restaurantGrpcHandler := grpcHandler.NewRestaurantGrpcHandler(restaurantService)

	proto.RegisterRestaurantServiceServer(grpcServer, restaurantGrpcHandler)

	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("Failed to listen on port 9090: %v", err)
	}

	log.Println("gRPC server is running on port 9091")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
