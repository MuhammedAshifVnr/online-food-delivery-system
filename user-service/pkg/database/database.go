package database

import (
	"log"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=0000 dbname=micro port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return DB
}
