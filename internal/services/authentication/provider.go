package authentication

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
)

type Auth struct {
	usersDirectory users.Querier
}

func New(usersDirectory users.Querier) *Auth {
	return &Auth{
		usersDirectory: usersDirectory,
	}
}
