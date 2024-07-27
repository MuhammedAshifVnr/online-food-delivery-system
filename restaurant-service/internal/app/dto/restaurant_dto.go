package dto

type RestaurantDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
}

type MenuItemDTO struct {
	ID           uint    `json:"id"`
	RestaurantID uint    `json:"restaurant_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Status       string  `json:"status"`
}
