package service

import (
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/app/dto"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/domain/model"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/restaurant-service/internal/repository"
)

// type RestaurantService interface {
// 	CreateRestaurant(dto dto.RestaurantDTO) (model.Restaurant, error)
// 	GetRestaurantByID(id uint) (model.Restaurant, error)
// 	CreateMenuItem(dto dto.MenuItemDTO) (model.MenuItem, error)
// 	GetMenuItemsByRestaurantID(restaurantID uint) ([]model.MenuItem, error)
// }

type RestaurantService struct {
	repository repository.RestaurantRepository
}

func NewRestaurantService(repository repository.RestaurantRepository) *RestaurantService {
	return &RestaurantService{repository}
}

func (s *RestaurantService) CreateRestaurant(dto dto.RestaurantDTO) (model.Restaurant, error) {
	restaurant := model.Restaurant{
		Name:        dto.Name,
		Description: dto.Description,
		Address:     dto.Address,
		Phone:       dto.Phone,
	}
	return s.repository.CreateRestaurant(restaurant)
}

func (s *RestaurantService) GetRestaurantByID() ([]model.Restaurant, error) {
	return s.repository.GetRestaurantByID()
}

func (s *RestaurantService) CreateMenuItem(dto dto.MenuItemDTO) (model.MenuItem, error) {
	menuItem := model.MenuItem{
		RestaurantID: dto.RestaurantID,
		Name:         dto.Name,
		Description:  dto.Description,
		Price:        dto.Price,
		Status:       dto.Status,
	}
	return s.repository.CreateMenuItem(menuItem)
}

func (s *RestaurantService) GetMenuItemsByRestaurantID(restaurantID uint) ([]model.MenuItem, error) {
	return s.repository.GetMenuItemsByRestaurantID(restaurantID)
}

func (s *RestaurantService) GetMenuItemByID(menuID uint)(model.MenuItem,error){
	return s.repository.GetMenuItemByID(menuID)
}