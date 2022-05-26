package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	iamV1 "github.com/mohammadVatandoost/ingbusiness/api/services/iam/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/services/iam"
	"github.com/mohammadVatandoost/ingbusiness/pkg/jwt"
	"net/http"
	"time"
)

func (s *Server) AddIngPage(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	var data iamV1.IngPage

	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for AddIngPage, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	userInfo, _ := c.Get(UserContextKey)

	ing, err := s.iamService.AddIngPage(ctx, &data, userInfo.(jwt.Message).UserID)
	if err != nil {
		s.logger.Errorf("can not get json data for AddIngPage, err: %s \n", err.Error())
		ErrorResponse(c, fmt.Errorf("%s", iam.ErrorCanNotAddIngPage).Error())
		return
	}

	APIResponse(c, http.StatusOK, nil, &iamV1.IngPage{
		OrganizationId: ing.OrganizationID,
		Name:           ing.Name,
		Id:             ing.ID,
	}, nil, "", nil)
}

func (s *Server) GetIngPages(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	var data iamV1.IngPage

	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for GetIngPages, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	ingPages, err := s.iamService.GetIngPageByOrganizationID(ctx, data.OrganizationId)
	if err != nil {
		s.logger.Errorf("can not get GetIngPageByOrganizationID, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	APIResponse(c, http.StatusOK, nil, ingPages, nil, "", nil)
}

func (s *Server) DeleteIngPage(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	var data iamV1.IngPage

	err := c.BindJSON(&data)
	if err != nil {
		s.logger.Errorf("can not get json data for DeleteIngPage, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	_, err = s.iamService.DeleteIngPage(ctx, data.Id)
	if err != nil {
		s.logger.Errorf("can not DeleteIngPage, err: %s \n", err.Error())
		ErrorResponse(c, err.Error())
		return
	}

	OKResponse(c, iam.DeletePage)
}
