package rest

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group("/v1")
	groupV1.GET("/experiment/disable/:id", s.GetDirectMessages)
	groupV1.GET("/experiment/enable/:id", s.SendDirectMessage)
}
