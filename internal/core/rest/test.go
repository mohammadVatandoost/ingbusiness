package rest

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	admin1_email = "admin1@my-domain.com"
	admin2_email = "admin2@my-domain.com"
)

func (s *Server) CleanTestData(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.TimeOut)*time.Second)
	defer cancel()

	err := s.userService.DeleteUserByEmail(ctx, admin1_email)
	if err != nil {
		ErrorResponse(c, err.Error())
		return
	}

	err = s.userService.DeleteUserByEmail(ctx, admin2_email)
	if err != nil {
		ErrorResponse(c, err.Error())
		return
	}

	OKResponse(c, "successfully cleaned test data")
}
