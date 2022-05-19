package authentication

import (
	"context"
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/authentication/v1"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
)

func (s *Service) SignUp(ctx context.Context, in *v1.SignUpRequest) (string, error) {

	pass, err := HashPassword(in.Password)
	if err != nil {
		return "", err
	}

	_, err = s.usersDirectory.AddUser(ctx,
		users.AddUserParams{
			Username: in.Username,
			Phone:    in.PhoneNumber,
			Email:    in.Email,
			Password: pass,
		})
	if err != nil {
		return "", err
	}

	return ConfirmEmail, nil
}
