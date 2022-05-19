package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (s *Server) SignUp(c *gin.Context) {

}

func (s *Server) SignIn(c *gin.Context) {

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
	// ToDo set to auth token
}
