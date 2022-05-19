package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	AuthTokenHeaderKey = "X-JWT-TOKEN"
	UserIDHeaderKey    = "X-USER-ID"
	UserDashboardPath  = "/dashboard"
)

func ErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"errors": gin.H{
			"message": message,
		},
	})
}

func SetAuthToken(c *gin.Context, token string) {
	c.Header(AuthTokenHeaderKey, token)
}

func SetUserID(c *gin.Context, id int32) {
	c.Header(UserIDHeaderKey, strconv.Itoa(id))
}

func APIResponse(c *gin.Context, status int, message []string, payload interface{}, err []string) {
	c.JSON(status, gin.H{
		"messages": message,
		"errors":   err,
		"payload":  payload,
	})
}
