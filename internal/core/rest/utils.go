package rest

import (
	"encoding/json"
	client "github.com/mohammadVatandoost/ingbusiness/api/services/client/v1"
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/notification/v1"
	"github.com/mohammadVatandoost/ingbusiness/pkg/notification"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	AuthTokenHeaderKey = "X-JWT-TOKEN"
	UserIDHeaderKey    = "X-USER-ID"
)

func ErrorResponse(c *gin.Context, message string) {
	//c.JSON(http.StatusBadRequest, gin.H{
	//	"error": message,
	//})

	APIResponse(c, http.StatusBadRequest, nil, nil,
		nil, "", []string{message})
}

func OKResponse(c *gin.Context, message string) {
	APIResponse(c, http.StatusOK, nil, nil,
		[]*v1.Notification{notification.MakeSuccess("", message)}, "", nil)
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

	jsonConfigs, err := json.Marshal(&res)
	if err != nil {
		logrus.Errorf("can not convert client response message to json, res: %s, err: %s \n",
			res.String(), err.Error())
	}

	//c.Data(status, gin.MIMEJSON, bytes)
	//c.String(status, string(bytes))
	c.JSON(status, gin.H{
		"metadata": string(jsonConfigs),
		"payload":  payload,
	})
}
