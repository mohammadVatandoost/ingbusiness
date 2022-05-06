package rest

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group("/v1")

	// authentication
	groupV1.POST("/signUp", s.SignUp)
	groupV1.POST("/signIn", s.SignIn)

	groupV1.POST("/getDirectMessages", s.GetDirectMessages)
	groupV1.POST("/sendDirectMessage", s.SendDirectMessage)
}
