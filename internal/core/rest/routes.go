package rest

const (
	versionPrefix = "/v1"
	authPrefix    = "/auth"
	userPrefix    = "/user"

	UserDashboardPath = versionPrefix + userPrefix + "/dashboard"

	AuthSignIn  = versionPrefix + authPrefix + "/signIn"
	AuthConfirm = versionPrefix + authPrefix + "/confirm"
)

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group(versionPrefix)

	authenticationGroup := groupV1.Group(authPrefix)

	// authentication
	authenticationGroup.POST("/signUp", s.SignUp)
	authenticationGroup.POST("/signIn", s.SignIn)
	authenticationGroup.GET("/{provider}/callback", s.OAuth2CallBack)
	authenticationGroup.GET("/{provider}", s.OAuth2)

	userGroup := groupV1.Group(userPrefix)
	userGroup.POST("/getDirectMessages", s.GetDirectMessages)
	userGroup.POST("/sendDirectMessage", s.SendDirectMessage)
}
