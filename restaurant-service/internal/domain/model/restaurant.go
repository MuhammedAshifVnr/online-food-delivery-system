package model

import "time"

type Restaurant struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Description string `gorm:"size:255"`
	Address     string `gorm:"size:255;not null"`
	Phone       string `gorm:"size:255;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MenuItem struct {
	ID           uint    `gorm:"primaryKey"`
	RestaurantID uint    `gorm:"not null"`
	Restaurant Restaurant
	Name         string  `gorm:"size:255;not null"`
	Description  string  `gorm:"size:255"`
	Price        float64 `gorm:"not null"`
	Status       string  `grom:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
