package dto

import "github.com/lib/pq"

// type OrderDTO struct {
// 	ID         uint    `json:"id"`
// 	UserID     uint    `json:"user_id"`
// 	Status     string  `json:"status"`
// 	TotalPrice float64 `json:"total_price"`
// }

type CreateOrderRequest struct {
	UserID  uint          `json:"user_id"`
	ItemIDs pq.Int64Array `json:"items_id" gorm:"type:integer[]"`
}

type UpdateOrderStatusRequest struct {
	OrderID uint   `json:"order_id"`
	Status  string `json:"status"`
}
