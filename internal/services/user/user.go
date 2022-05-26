package user

import (
	"context"
)

func (s *Service) DeleteUserByEmail(ctx context.Context, email string) error {
	u, err := s.usersDirectory.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	_, err = s.usersDirectory.DeleteUser(ctx, u.ID)
	return err
}
