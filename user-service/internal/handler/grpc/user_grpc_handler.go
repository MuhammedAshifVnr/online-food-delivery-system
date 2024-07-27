package grpc

import (
	"context"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/proto"
)

type UserGRPCHandler struct {
	userService service.UserService
	proto.UnimplementedUserServiceServer
}

func NewUserGRPCHandler(userService service.UserService) *UserGRPCHandler {
	return &UserGRPCHandler{
		userService: userService,
	}
}

func (h *UserGRPCHandler) GetUserDetails(ctx context.Context, req *proto.GetUserDetailsRequest) (*proto.UserResponse, error) {
	_, err := h.userService.GetUserByID(uint(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &proto.UserResponse{
		Exists: true,
	}, nil
}
