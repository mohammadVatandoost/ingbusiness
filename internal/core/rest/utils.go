package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": gin.H{
			"message": message,
		},
	})
}

func APIResponse(c *gin.Context, status int, message []string, payload interface{}, err []string) {
	c.JSON(status, gin.H{
		"messages": message,
		"errors":   err,
		"payload":  payload,
	})
}
