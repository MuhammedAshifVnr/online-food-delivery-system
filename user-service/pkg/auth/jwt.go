package auth

import (
	"errors"
	"time"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/internal/domain/model"
	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/pkg/logger"
	"github.com/golang-jwt/jwt"
)

var (
	jwtSecret             = []byte("my_jwt_secret")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(user model.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		logger.LogError("Failed to generate JWT", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
