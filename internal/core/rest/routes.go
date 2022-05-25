package rest

const (
	versionPrefix      = "/v1"
	authPrefix         = "/auth"
	userPrefix         = "/user"
	organizationPrefix = "/organization"

	UserDashboardPath = versionPrefix + userPrefix + "/dashboard"

	AuthSignIn  = versionPrefix + authPrefix + "/signIn"
	AuthConfirm = versionPrefix + authPrefix + "/confirm"
)

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group(versionPrefix)

	// authentication
	authenticationGroup := groupV1.Group(authPrefix)
	authenticationGroup.POST("/signUp", s.SignUp)
	authenticationGroup.POST("/signIn", s.SignIn)
	authenticationGroup.GET("/{provider}/callback", s.OAuth2CallBack)
	authenticationGroup.GET("/{provider}", s.OAuth2)

	// user
	userGroup := groupV1.Group(userPrefix, s.authMiddleware())
	userGroup.POST("/addOrganization", s.AddOrganization)
	userGroup.POST("/getUserProfile", s.GetUserProfile)

	// organization permission
	organizationGroup := userGroup.Group(organizationPrefix, s.iamMiddleware())
	organizationGroup.POST("/giveOrganizationPermissionByEmail", s.GiveOrganizationPermissionByEmail)
	organizationGroup.POST("/deleteOrganizationPermissionByEmail", s.DeleteOrganizationPermissionByEmail)
	organizationGroup.POST("/deleteOrganization", s.DeleteOrganization)

	// frequent messages
	organizationGroup.POST("/saveMessage", s.SaveMessage)
	organizationGroup.POST("/getSavedMessages", s.GetSavedMessages)
	organizationGroup.POST("/deleteSavedMessage", s.DeleteSavedMessage)
	organizationGroup.POST("/updateSavedMessage", s.UpdateSavedMessage)

	// messenger
	//organizationGroup.POST("/saveMessage", s.SaveMessage)
	//organizationGroup.POST("/getSavedMessages", s.GetSavedMessages)
	//organizationGroup.POST("/deleteSavedMessage", s.DeleteSavedMessage)
	//organizationGroup.POST("/updateSavedMessage", s.UpdateSavedMessage)

}
