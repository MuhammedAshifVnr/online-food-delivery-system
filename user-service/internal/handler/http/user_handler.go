package http

import (
	"net/http"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/app/dto"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/app/service"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/repository"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/pkg/database"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindBodyWithJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := h.userService.RegisterUser(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var user dto.UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.userService.LoginUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetUserDetails(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userDTO.ID = c.GetUint("userID")
	updatedUser, err := h.userService.UpdateUser(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func RegisterUserRoutes(router *gin.Engine) {
	userHandler := NewUserHandler(service.NewUserService(repository.NewUserRepository(database.GetDB())))

	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.GET("/me", middleware.AuthMiddleware(), userHandler.GetUserDetails)
		userRoutes.PUT("/me", middleware.AuthMiddleware(), userHandler.UpdateProfile)
	}
}
