package rest

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group("/v1")

	// authentication
	groupV1.POST("/auth/signUp", s.SignUp)
	groupV1.POST("/auth/signIn", s.SignIn)
	groupV1.GET("/auth/{provider}/callback", s.OAuth2CallBack)
	groupV1.GET("/auth/{provider}", s.OAuth2)

	groupV1.POST("/getDirectMessages", s.GetDirectMessages)
	groupV1.POST("/sendDirectMessage", s.SendDirectMessage)
}
