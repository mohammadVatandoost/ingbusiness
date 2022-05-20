package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/authentication"
	"github.com/mohammadVatandoost/ingbusiness/pkg/jwt"
	"github.com/mohammadVatandoost/ingbusiness/pkg/notification"
	"net/http"
	"time"
)

func (s *Server) AddOrganization(c *gin.Context) {
	//s.iamService.AddOrganization()
}

func (s *Server) GiveOrganizationPermissionByEmail(c *gin.Context) {
	//s.iamService.GivePermissionByEmail()
}

func (s *Server) DeleteOrganizationPermissionByEmail(c *gin.Context) {
	//s.iamService.DeletePermissionByEmail()
}

func (s *Server) DeleteOrganization(c *gin.Context) {
	//ToDo: just Owner
	//s.iamService.DeleteOrganization()
}

func (s *Server) iamMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userInfo, exists := c.Get(UserContextKey)
		if !exists {
			s.logger.Warn("user is not login \n")
			APIResponse(c, http.StatusBadRequest, nil, nil,
				notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
				AuthSignIn, nil)
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
		defer cancel()

		organizationName := c.GetHeader(OrganizationNameHeaderKey)
		u, ok := userInfo.(jwt.Message)
		if !ok {
			s.logger.Warn("user info is not set correctly \n")
			APIResponse(c, http.StatusBadRequest, nil, nil,
				notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
				AuthSignIn, nil)
			c.Abort()
			return
		}
		roleType, err := s.iamService.GetUserRoleInOrganization(ctx, u.UserID, organizationName)
		if err != nil {
			s.logger.Warnf("user does not have any role, userID: %v, err: %s \n",
				u.UserID, err.Error())
			APIResponse(c, http.StatusBadRequest, nil, nil,
				notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
				AuthSignIn, nil)
			c.Abort()
			return
		}

		c.Set(UserRoleContextKey, roleType)
		c.Next()
	}
}
