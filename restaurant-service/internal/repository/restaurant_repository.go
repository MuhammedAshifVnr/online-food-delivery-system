package repository

import (
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/domain/model"
	"gorm.io/gorm"
)

type RestaurantRepository interface {
	CreateRestaurant(restaurant model.Restaurant) (model.Restaurant, error)
	GetRestaurantByID() ([]model.Restaurant, error)
	CreateMenuItem(menuItem model.MenuItem) (model.MenuItem, error)
	GetMenuItemsByRestaurantID(restaurantID uint) ([]model.MenuItem, error)
	GetMenuItemByID(itemID uint) (model.MenuItem, error)
}

type restaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &restaurantRepository{db}
}

func (r *restaurantRepository) CreateRestaurant(restaurant model.Restaurant) (model.Restaurant, error) {
	if err := r.db.Create(&restaurant).Error; err != nil {
		return model.Restaurant{}, err
	}
	return restaurant, nil
}

func (r *restaurantRepository) GetRestaurantByID() ([]model.Restaurant, error) {
	var restaurant []model.Restaurant
	if err := r.db.Find(&restaurant).Error; err != nil {
		return []model.Restaurant{}, err
	}
	return restaurant, nil
}

func (r *restaurantRepository) CreateMenuItem(menuItem model.MenuItem) (model.MenuItem, error) {
	if err := r.db.Create(&menuItem).Error; err != nil {
		return model.MenuItem{}, err
	}
	return menuItem, nil
}

func (r *restaurantRepository) GetMenuItemsByRestaurantID(restaurantID uint) ([]model.MenuItem, error) {
	var menuItems []model.MenuItem
	if err := r.db.Where("restaurant_id = ?", restaurantID).Find(&menuItems).Error; err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (r *restaurantRepository) GetMenuItemByID(itemID uint) (model.MenuItem, error) {
	var menu model.MenuItem
	if err := r.db.Where("id = ?", itemID).First(&menu).Error; err != nil {
		return model.MenuItem{}, err
	}
	return menu, nil
}
