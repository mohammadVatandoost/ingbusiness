package rest

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group("/v1")

	authenticationGroup := groupV1.Group("/auth")

	// authentication
	authenticationGroup.POST("/signUp", s.SignUp)
	authenticationGroup.POST("/signIn", s.SignIn)
	authenticationGroup.GET("/{provider}/callback", s.OAuth2CallBack)
	authenticationGroup.GET("/{provider}", s.OAuth2)

	userGroup := groupV1.Group("/user")
	userGroup.POST("/getDirectMessages", s.GetDirectMessages)
	userGroup.POST("/sendDirectMessage", s.SendDirectMessage)
}
