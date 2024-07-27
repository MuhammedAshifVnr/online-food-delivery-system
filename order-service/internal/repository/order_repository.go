package repository

import (
	"fmt"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/order-service/internal/domain/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order model.Order) (model.Order, error)
	UpdateStatus(orderID uint, status string) error
	GetByID(orderID uint) (model.Order, error)
	GetByUserID(userID uint) ([]model.Order, error)
	GetAll() ([]model.Order, error)
	CompletePayment(orderId string) (bool, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order model.Order) (model.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *orderRepository) UpdateStatus(orderID uint, status string) error {
	return r.db.Model(&model.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *orderRepository) GetByID(orderID uint) (model.Order, error) {
	var order model.Order
	err := r.db.First(&order, orderID).Error
	return order, err
}

func (r *orderRepository) GetByUserID(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) CompletePayment(orderId string) (bool, error) {
	var order model.Order
	fmt.Println("-----", orderId)
	if err := r.db.Where("payment_id = ?", orderId).First(&order).Error; err != nil {
		fmt.Println("err===",err)
		return false, err
	}
	order.Status = "Success"
	if err := r.db.Save(&order).Error; err != nil {
		return false, err
	}
	return true, nil

}
