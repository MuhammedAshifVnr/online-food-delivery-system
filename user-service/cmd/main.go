package main

import (
	"log"
	"net"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/handler/http"
	grpcHandler "github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/handler/grpc"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/repository"
	postgres "github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/pkg/database"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	postgres.Init()
	router := gin.Default()
	http.RegisterUserRoutes(router)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	db:=postgres.GetDB()
	repo:=repository.NewUserRepository(db)
	grpcService:=service.NewUserService(repo)
	userGrpcHandler:=grpcHandler.NewUserGRPCHandler(grpcService)
	grpcServer:=grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userGrpcHandler)

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("Failed to listen on port 9090: %v", err)
	}

	log.Println("gRPC server is running on port 9090")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
