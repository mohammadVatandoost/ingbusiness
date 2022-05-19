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
		"error": message,
	})
}

func SetAuthToken(c *gin.Context, token string) {
	c.Header(AuthTokenHeaderKey, token)
}

func SetUserID(c *gin.Context, id int32) {
	c.Header(UserIDHeaderKey, strconv.Itoa(int(id)))
}

func APIResponse(c *gin.Context, status int, message []string, payload interface{},
	notifications interface{}, redirectURL string, err []string) {
	c.JSON(status, gin.H{
		"messages":      message,
		"notifications": notifications,
		"payload":       payload,
		"redirect":      redirectURL,
		"errors":        err,
	})
}
