package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderID    uint          `json:"order_id" gorm:"unique"`
	UserID     uint          `json:"user_id"`
	ItemsID    pq.Int64Array `json:"items_id" gorm:"type:integer[]"`
	Status     string        `json:"status"`
	TotalPrice float64       `json:"total_price"`
	PaymentID  string
}
