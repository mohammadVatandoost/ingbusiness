package iam

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/access"
	"github.com/mohammadVatandoost/ingbusiness/internal/ingpages"
	"github.com/mohammadVatandoost/ingbusiness/internal/organization"
	roles "github.com/mohammadVatandoost/ingbusiness/internal/role"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
)

type Service struct {
	organizationDirectory organization.Querier
	rolesDirectory        roles.Querier
	accessDirectory       access.Querier
	usersDirectory        users.Querier
	ingDirectory          ingpages.Querier
}

func New(organizationDirectory organization.Querier, rolesDirectory roles.Querier,
	ingDirectory ingpages.Querier,
	accessDirectory access.Querier, usersDirectory users.Querier) *Service {
	return &Service{
		organizationDirectory: organizationDirectory,
		rolesDirectory:        rolesDirectory,
		accessDirectory:       accessDirectory,
		usersDirectory:        usersDirectory,
		ingDirectory:          ingDirectory,
	}
}
