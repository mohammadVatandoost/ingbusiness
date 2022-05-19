package iam

import (
	"github.com/mohammadVatandoost/ingbusiness/internal/organization"
)

type Service struct {
	organizationDirectory organization.Querier
}

func New(organizationDirectory organization.Querier) *Service {
	return &Service{
		organizationDirectory: organizationDirectory,
	}
}
