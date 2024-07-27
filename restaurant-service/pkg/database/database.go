package database

import (
	"log"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=0000 dbname=restaurant port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB.AutoMigrate(&model.MenuItem{},&model.Restaurant{})
}

func GetDB() *gorm.DB {
	return DB
}
