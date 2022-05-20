package iam

import (
	"context"
	"github.com/mohammadVatandoost/ingbusiness/internal/access"
	"github.com/mohammadVatandoost/ingbusiness/internal/organization"
)

func (s *Service) AddOrganization(ctx context.Context,
	params organization.AddOrganizationParams) (*organization.Organization, error) {
	Organization, err := s.organizationDirectory.AddOrganization(ctx, params)
	if err != nil {
		return nil, err
	}
	_, err = s.accessDirectory.AddAccess(ctx, access.AddAccessParams{
		UserID:         params.OwnerID,
		OrganizationID: Organization.ID,
		RoleID:         OwnerRole,
	})
	if err != nil {
		_, _ = s.organizationDirectory.DeleteOrganization(ctx, Organization.ID)
		return nil, err
	}
	return &Organization, nil
}

func (s *Service) DeleteOrganization(ctx context.Context, id int32) error {
	_, err := s.organizationDirectory.DeleteOrganization(ctx, id)
	if err != nil {
		return err
	}
	_, err = s.accessDirectory.DeleteAccessByOrganizationID(ctx, id)
	return err
}
