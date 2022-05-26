package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	iamV1 "github.com/mohammadVatandoost/ingbusiness/api/services/iam/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/organization"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/authentication"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/iam"
	"github.com/mohammadVatandoost/ingbusiness/pkg/jwt"
	"github.com/mohammadVatandoost/ingbusiness/pkg/notification"
	"net/http"
	"time"
)

func (s *Server) AddOrganization(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	var data iamV1.Organization
	userInfo, _ := c.Get(UserContextKey)

	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for AddOrganization, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	organ, err := s.iamService.AddOrganization(ctx, organization.AddOrganizationParams{
		OwnerID: userInfo.(jwt.Message).UserID,
		Name:    data.Name,
	})
	if err != nil {
		s.logger.Errorf("can not add organ, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	APIResponse(c, http.StatusOK, nil,
		iamV1.Organization{Id: organ.ID, Name: organ.Name, OwnerId: organ.OwnerID},
		nil,
		"", nil)
}

func (s *Server) GetOrganizationUserHasRole(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	userInfo, _ := c.Get(UserContextKey)

	access, err := s.iamService.GetOrganizationUserHasRole(ctx, userInfo.(jwt.Message).UserID)
	if err != nil {
		s.logger.Errorf("can not GetOrganizationUserHasRole, id: %v, err: %s \n",
			userInfo.(jwt.Message).UserID, err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	var res []iamV1.OrganizationAccess
	for _, a := range access {
		res = append(res, iamV1.OrganizationAccess{
			OrganizationID:   a.OrganizationID,
			OrganizationName: a.OrganizationName,
			RoleID:           a.RoleID,
			UserID:           a.UserID,
		})
	}
	APIResponse(c, http.StatusOK, nil,
		res,
		nil,
		"", nil)
}

func (s *Server) GiveOrganizationPermissionByEmail(c *gin.Context) {
	//just Owner
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()
	var data iamV1.OrganizationPermission
	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for GiveOrganizationPermissionByEmail, "+
			"ID: %v, Email: %v, err: %s \n",
			data.Id, data.Email, err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	userInfo, _ := c.Get(UserContextKey)
	roleType, ok := c.Get(UserRoleContextKey)
	if !ok {
		s.logger.Warnf("user role is not set correctly \n")
		APIResponse(c, http.StatusBadRequest, nil, nil,
			notification.BuildNotification(notification.MakeError("", authentication.ErrorNotOwner)),
			"", nil)
		return
	}

	if roleType.(int32) != iam.OwnerRole {
		s.logger.Warnf("user role is not Owner , userID: %v, roleType: %v \n",
			userInfo.(jwt.Message).UserID, roleType)
		APIResponse(c, http.StatusBadRequest, nil, nil,
			notification.BuildNotification(notification.MakeError("", authentication.ErrorNotOwner)),
			"", nil)
		return
	}

	err = s.iamService.GivePermissionByEmail(ctx, data.Email, data.Id)
	if err != nil {
		s.logger.Errorf("can not GivePermissionByEmail, Email: %s, organID: %v, err: %s \n",
			data.Email, data.Id, err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	APIResponse(c, http.StatusOK, nil, nil, nil,
		"", nil)
}

func (s *Server) DeleteOrganizationPermissionByEmail(c *gin.Context) {
	//just Owner
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()
	var data iamV1.OrganizationPermission
	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for GiveOrganizationPermissionByEmail, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	roleType, ok := c.Get(UserRoleContextKey)
	if !ok {
		s.logger.Warn("user role is not set correctly \n")
		APIResponse(c, http.StatusBadRequest, nil, nil,
			notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
			AuthSignIn, nil)
		return
	}

	if roleType != iam.OwnerRole {
		s.logger.Warn("user role is not Owner \n")
		APIResponse(c, http.StatusBadRequest, nil, nil,
			notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
			AuthSignIn, nil)
		return
	}

	err = s.iamService.DeletePermissionByEmail(ctx, data.Email, data.Id)
	if err != nil {
		s.logger.Errorf("can not DeletePermissionByEmail, Email: %s, organID: %v, err: %s \n",
			data.Email, data.Id, err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	APIResponse(c, http.StatusOK, nil, nil, nil,
		"", nil)
}

func (s *Server) DeleteOrganization(c *gin.Context) {
	//just Owner
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	var data iamV1.Organization

	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for DeleteOrganization, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	roleType, ok := c.Get(UserRoleContextKey)
	if !ok {
		s.logger.Warn("user role is not set correctly \n")
		APIResponse(c, http.StatusBadRequest, nil, nil,
			notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
			AuthSignIn, nil)
		return
	}

	if roleType != iam.OwnerRole {
		s.logger.Warn("user role is not Owner \n")
		APIResponse(c, http.StatusBadRequest, nil, nil,
			notification.BuildNotification(notification.MakeError("", authentication.ErrorNotAuthorized)),
			AuthSignIn, nil)
		return
	}

	err = s.iamService.DeleteOrganization(ctx, data.Id)
	if err != nil {
		s.logger.Errorf("can not DeleteOrganization, id:%v.  err: %s \n", data.Id, err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	APIResponse(c, http.StatusOK, nil, nil, nil,
		"", nil)
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
