package rest

import (
	client "github.com/mohammadVatandoost/ingbusiness/api/services/client/v1"
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/notification/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	AuthTokenHeaderKey = "X-JWT-TOKEN"
	UserIDHeaderKey    = "X-USER-ID"
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
	notifications []*v1.Notification, redirectURL string, errors []string) {
	res := client.Response{
		Notifications: notifications,
		Redirect:      redirectURL,
		Errors:        errors,
		Messages:      message,
	}
	bytes, err := proto.Marshal(&res)
	if err != nil {
		logrus.Errorf("can not convert client response message to json, res: %s, err: %s \n",
			res.String(), err.Error())
	}
	c.Data(status, "application/json", bytes)
	//c.String(status, string(bytes))
	//c.JSON(status, gin.H{
	//	"messages":      message,
	//	"notifications": notifications,
	//	"payload":       payload,
	//	"redirect":      redirectURL,
	//	"errors":        err,
	//})
}
