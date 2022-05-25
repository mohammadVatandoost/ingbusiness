package user

import (
	"context"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
)

func (s *Service) GetProfile(ctx context.Context, userID int32) (users.User, error) {
	return s.usersDirectory.GetUser(ctx, userID)
}
