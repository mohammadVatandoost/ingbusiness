package rest

const (
	versionPrefix      = "/v1"
	authPrefix         = "/auth"
	userPrefix         = "/user"
	organizationPrefix = "/organization"
	testPrefix         = "/test"

	UserDashboardPath = versionPrefix + userPrefix + "/dashboard"

	AuthSignIn  = versionPrefix + authPrefix + "/signIn"
	AuthConfirm = versionPrefix + authPrefix + "/confirm"
)

func (s *Server) Routes() {
	router := s.Engine
	groupV1 := router.Group(versionPrefix)

	//test
	testGroup := groupV1.Group(testPrefix)
	testGroup.POST("/cleanTestData", s.CleanTestData)
	// authentication
	authenticationGroup := groupV1.Group(authPrefix)
	authenticationGroup.POST("/signUp", s.SignUp)
	authenticationGroup.POST("/signIn", s.SignIn)
	authenticationGroup.GET("/{provider}/callback", s.OAuth2CallBack)  // for test: http://localhost:9077/v1/auth/facebook/callback
	authenticationGroup.GET("/{provider}", s.OAuth2)

	// user
	userGroup := groupV1.Group(userPrefix, s.authMiddleware())
	userGroup.POST("/addOrganization", s.AddOrganization)
	userGroup.POST("/getOrganizationUserHasRole", s.GetOrganizationUserHasRole)
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

	// ing pages
	organizationGroup.POST("/addIngPage", s.AddIngPage)
	organizationGroup.POST("/getIngPages", s.GetIngPages)
	organizationGroup.POST("/deleteIngPage", s.DeleteIngPage)
	organizationGroup.POST("/updateIngPage", s.UpdateIngPage)

}
