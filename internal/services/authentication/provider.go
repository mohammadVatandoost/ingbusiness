package authentication

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
)

type Service struct {
	usersDirectory users.Querier
	conf           Config
}

func New(usersDirectory users.Querier, conf Config) *Service {
	return &Service{
		usersDirectory: usersDirectory,
		conf:           conf,
	}
}
