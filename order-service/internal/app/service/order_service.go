package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"strconv"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/app/dto"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/domain/model"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/repository"
	proto "github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/proto"
)

type OrderService struct {
	repo       repository.OrderRepository
	userClient proto.UserServiceClient
	restClient proto.RestaurantServiceClient
	payClient  proto.PaymentServiceClient
}

func NewOrderService(repo repository.OrderRepository, user proto.UserServiceClient, rest proto.RestaurantServiceClient, pay proto.PaymentServiceClient) *OrderService {
	return &OrderService{repo: repo, userClient: user, restClient: rest, payClient: pay}
}

func (s *OrderService) PlaceOrder(req dto.CreateOrderRequest) (model.Order, error) {
	userReq := &proto.GetUserDetailsRequest{Id: uint32(req.UserID)}
	userResp, err := s.userClient.GetUserDetails(context.Background(), userReq)
	if err != nil {
		return model.Order{}, err
	}

	if !userResp.Exists {
		return model.Order{}, errors.New("user does not exist")
	}
	var totalAmoutn uint
	for _, item := range req.ItemIDs {
		restReq := &proto.GetMenuItemRequest{Id: uint32(item)}
		restRes, err := s.restClient.GetMenuItem(context.Background(), restReq)
		if err != nil {
			return model.Order{}, err
		}
		if restRes.Status == "u" {
			return model.Order{}, fmt.Errorf("item %v is unavailable", restRes.Name)
		}
		totalAmoutn += uint(restRes.Price)
	}
	orderID, _ := strconv.ParseUint(generateRandomNumber(), 10, 64)
	payReq := &proto.NewOrderRequest{Price: uint32(totalAmoutn), OrderId: uint32(orderID)}
	payRes, erro := s.payClient.NewOrder(context.Background(), payReq)
	if erro != nil {
		fmt.Println("payment",erro)
		return model.Order{}, erro
	}

	order := model.Order{
		OrderID:    uint(orderID),
		TotalPrice: float64(totalAmoutn),
		UserID:     req.UserID,
		ItemsID:    req.ItemIDs,
		Status:     "payment pending",
		PaymentID:  payRes.RazorOrderId,
	}
	return s.repo.Create(order)
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	return s.repo.UpdateStatus(orderID, status)
}

func (s *OrderService) GetOrderDetails(orderID uint) (model.Order, error) {
	return s.repo.GetByID(orderID)
}

func (s *OrderService) GetOrderHistory(userID uint) ([]model.Order, error) {
	return s.repo.GetByUserID(userID)
}

func (s *OrderService) GetAllOrders() ([]model.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) CompletePayment(orderID string) (bool, error) {
	fmt.Println("-id====",orderID)
	return s.repo.CompletePayment(orderID)

}

func generateRandomNumber() string {
	const charset = "123456789"
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(err)
	}
	for i, b := range randomBytes {
		randomBytes[i] = charset[b%byte(len(charset))]
	}
	return string(randomBytes)
}
