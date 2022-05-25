package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	userV1 "github.com/mohammadVatandoost/ingbusiness/api/services/user/v1"
	"github.com/mohammadVatandoost/ingbusiness/pkg/jwt"
	"net/http"
	"time"
)

func (s *Server) GetUserProfile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	userInfo, _ := c.Get(UserContextKey)
	u, err := s.userService.GetProfile(ctx, userInfo.(jwt.Message).UserID)
	if err != nil {
		s.logger.Errorf("can not GetProfile, userID: %v, err: %s \n",
			userInfo.(jwt.Message).UserID, err.Error())
		ErrorResponse(c, err.Error())
		return
	}
	APIResponse(c, http.StatusOK, nil,
		userV1.Profile{
			Id:           u.ID,
			Name:         u.Name,
			Email:        u.Email,
			Phone:        u.Phone,
			Username:     u.Username,
			ProfileImage: u.ProfileImage,
		},
		nil,
		"", nil)

}
