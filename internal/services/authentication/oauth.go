package authentication

import (
	"context"
	"github.com/markbates/goth"
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/authentication/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
	"github.com/mohammadVatandoost/ingbusiness/pkg/jwt"
	"time"
)

//oauth2 :https://www.loginradius.com/blog/engineering/google-authentication-with-golang-and-goth/
// https://github.com/markbates/goth
func (s *Service) OAuth2(ctx context.Context, userOAuth goth.User) (*v1.SignInResponse, error) {
	user, err := s.usersDirectory.GetUserByEmail(ctx, userOAuth.Email)
	if err != nil {
		s.logger.Infof("registering user by OAuth 2, user email: %s \n", userOAuth.Email)
		// ToDo: handle user name is not unique
		user, err = s.usersDirectory.AddUser(ctx,
			users.AddUserParams{
				Username: userOAuth.Name + userOAuth.UserID,
				Email:    userOAuth.Email,
			})
		if err != nil {
			return nil, err
		}
	}

	token, err := jwt.Sign(jwt.Message{
		UserID:    user.ID,
		Email:     user.Email,
		UserName:  user.Name,
		Timestamp: time.Now().Unix(),
	}, s.jwtSecret)

	if err != nil {
		return nil, err
	}

	return &v1.SignInResponse{
		Token:  token,
		UserID: user.ID,
	}, nil
}