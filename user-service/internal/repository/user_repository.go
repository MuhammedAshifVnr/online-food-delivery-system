package repository

import (
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(id uint) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
