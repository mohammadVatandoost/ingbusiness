package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/authentication/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/authentication"
	"github.com/mohammadVatandoost/ingbusiness/pkg/notification"
	"net/http"
	"time"
)

func (s *Server) SignUp(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	var data v1.SignUpRequest

	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for SignUp, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	message, err := s.authenticationService.SignUp(ctx, &data)
	if err != nil {
		s.logger.Errorf("can not SignUp, userData: %v, err: %s \n", data.String(), err.Error())
		ErrorResponse(c, err.Error())
		return
	}
	APIResponse(c, http.StatusOK, []string{message}, nil, nil, AuthConfirm, nil)
}

func (s *Server) SignIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	var data v1.SignInRequest

	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for SignIn, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	res, err := s.authenticationService.SignIn(ctx, &data)
	if err != nil {
		s.logger.Errorf("can not SignIn, userData: %v, err: %s \n", data.String(), err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	SetAuthToken(c, res.Token)
	//SetUserID(c, res.UserID)
	c.Status(http.StatusOK)
	//c.Redirect(http.StatusOK, UserDashboardPath)

}

func (s *Server) OAuth2(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (s *Server) OAuth2CallBack(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		s.logger.Errorf("can not get user profile by OAuth2, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	res, err := s.authenticationService.OAuth2(c, user)
	if err != nil {
		s.logger.Errorf("can not get user profile by OAuth2, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}
	SetAuthToken(c, res.Token)
	//SetUserID(c, res.UserID)
	c.Status(http.StatusOK)
	//c.Redirect(http.StatusOK, UserDashboardPath)
}

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(AuthTokenHeaderKey)
		userInfo, err := s.authenticationService.ValidateJWT(token)
		if err != nil {
			s.logger.Warnf("can not ValidateJWT user token, err: %s \n", err.Error())
			APIResponse(c, http.StatusBadRequest, nil, nil,
				notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
				AuthSignIn, nil)
			c.Abort()
			return
		}
		c.Set(UserContextKey, userInfo)
		c.Next()
	}
}
