package user

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
)

type Service struct {
	usersDirectory users.Querier
}

func New(usersDirectory users.Querier) *Service {
	return &Service{
		usersDirectory: usersDirectory,
	}
}
