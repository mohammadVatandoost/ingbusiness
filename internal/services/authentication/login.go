package authentication

import (
	"context"
	"fmt"
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/authentication/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
	"github.com/mohammadVatandoost/ingbusiness/pkg/jwt"
	"time"
)

func (s *Service) SignIn(ctx context.Context, in *v1.SignInRequest) (*v1.SignInResponse, error) {

	pass, err := HashPassword(in.Password)
	if err != nil {
		return nil, err
	}

	var user users.User
	if in.Email != "" {
		user, err = s.usersDirectory.GetUserByEmail(ctx, in.Email)
		if err != nil {
			return nil, err
		}
		if CheckPasswordHash(pass, user.Password) {
			return nil, fmt.Errorf("%v", ErrorWrongEmailOrPassword)
		}
	} else {
		user, err = s.usersDirectory.GetUserByUserName(ctx, in.Username)
		if err != nil {
			return nil, err
		}
		if CheckPasswordHash(pass, user.Password) {
			return nil, fmt.Errorf("%v", ErrorWrongUsernameOrPassword)
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
