package middleware

import (
	"net/http"

	"github.com/MuhammedAshifVnr/online-food-delivery-system/user-service/pkg/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
		}
		//tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims, err := auth.ValidateJWT(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}

}
