package response

import (
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, status int, data interface{}, message string) {
	c.JSON(status, gin.H{
		"data":    data,
		"message": message,
	})
}
