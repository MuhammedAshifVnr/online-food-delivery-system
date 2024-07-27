package service

import (
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/app/dto"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/domain/model"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/repository"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/pkg/auth"
)

type UserService interface {
	RegisterUser(userDTO dto.UserDTO) (dto.UserDTO, error)
	LoginUser(email, password string) (string, error)
	GetUserByID(id uint) (dto.UserDTO, error)
	UpdateUser(userDTO dto.UserDTO) (dto.UserDTO, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(userDTO dto.UserDTO) (dto.UserDTO, error) {
	user := model.User{
		ID:       userDTO.ID,
		Username: userDTO.Email,
		Email:    userDTO.Email,
		Password: userDTO.Password,
	}

	user.HashPassword()
	createUser, err := s.repo.CreateUser(user)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return dto.UserDTO{
		ID:       createUser.ID,
		Username: createUser.Username,
		Email:    createUser.Email,
	}, nil
}

func (s *userService) LoginUser(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if !user.CheckPassword(password) {
		return "", auth.ErrInvalidCredentials
	}
	token, err := auth.GenerateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) GetUserByID(id uint) (dto.UserDTO, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return dto.UserDTO{}, err
	}
	return dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *userService) UpdateUser(userDTO dto.UserDTO) (dto.UserDTO, error) {
	user := model.User{
		ID:       userDTO.ID,
		Username: userDTO.Username,
		Email:    userDTO.Email,
	}

	updatedUser, err := s.repo.UpdateUser(user)
	if err != nil {
		return dto.UserDTO{}, err
	}

	return dto.UserDTO{
		ID:       updatedUser.ID,
		Username: updatedUser.Username,
		Email:    updatedUser.Email,
	}, nil
}
