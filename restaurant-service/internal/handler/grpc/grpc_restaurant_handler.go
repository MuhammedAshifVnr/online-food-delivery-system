package grpc

import (
	"context"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/app/service"
	proto "github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/proto"
)

type RestaurantGrpcHandler struct {
	service *service.RestaurantService
	proto.UnimplementedRestaurantServiceServer
}

func NewRestaurantGrpcHandler(service *service.RestaurantService) *RestaurantGrpcHandler {
	return &RestaurantGrpcHandler{service: service}
}

// func (h *RestaurantGrpcHandler) AddRestaurant(ctx context.Context, req *proto.AddRestaurantRequest) (*proto.RestaurantResponse, error) {
// 	restaurant := dto.RestaurantDTO{
// 		Name:        req.Name,
// 		Description: req.Description,
// 		Address:     req.Address,
// 		Phone:       req.Phone,
// 	}

// 	createdRestaurant, err := h.service.CreateRestaurant(restaurant)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &proto.RestaurantResponse{
// 		Id:          uint32(createdRestaurant.ID),
// 		Name:        createdRestaurant.Name,
// 		Description: createdRestaurant.Description,
// 		Address:     createdRestaurant.Address,
// 		Phone:       createdRestaurant.Phone,
// 	}, nil
// }

// func (h *RestaurantGrpcHandler) GetRestaurant(ctx context.Context, req *proto.GetRestaurantRequest) (*proto.RestaurantResponse, error) {
// 	restaurant, err := h.service.GetRestaurantByID(uint(req.Id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &proto.RestaurantResponse{
// 		Id:          uint32(restaurant.ID),
// 		Name:        restaurant.Name,
// 		Description: restaurant.Description,
// 		Address:     restaurant.Address,
// 		Phone:       restaurant.Phone,
// 	}, nil
// }

func (h *RestaurantGrpcHandler) GetMenuItem(ctx context.Context, req *proto.GetMenuItemRequest) (*proto.MenuItemResponse, error) {
	menuItem, err := h.service.GetMenuItemByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &proto.MenuItemResponse{
		RestaurantId: uint32(menuItem.RestaurantID),
		Name:         menuItem.Name,
		Description:  menuItem.Description,
		Price:        float32(menuItem.Price),
		Status: menuItem.Status,
	}, nil
}
