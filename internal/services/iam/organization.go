package iam

import (
	"context"
	"github.com/mohammadVatandoost/ingbusiness/internal/access"
	"github.com/mohammadVatandoost/ingbusiness/internal/organization"
	roles "github.com/mohammadVatandoost/ingbusiness/internal/role"
)

func (s *Service) AddOrganization(ctx context.Context,
	params organization.AddOrganizationParams) (*organization.Organization, error) {
	Organization, err := s.organizationDirectory.AddOrganization(ctx, params)
	if err != nil {
		return nil, err
	}

	role, err := s.rolesDirectory.AddRole(ctx,
		roles.AddRoleParams{
			OrganizationID: Organization.ID,
			CreatorID:      params.OwnerID,
			RoleType:       OwnerRole,
		})
	if err != nil {
		return nil, err
	}
	_, err = s.rolesDirectory.AddRole(ctx,
		roles.AddRoleParams{
			OrganizationID: Organization.ID,
			CreatorID:      params.OwnerID,
			RoleType:       AdminRole,
		})
	if err != nil {
		return nil, err
	}
	_, err = s.accessDirectory.AddAccess(ctx, access.AddAccessParams{
		UserID:           params.OwnerID,
		OrganizationName: params.Name,
		OrganizationID:   Organization.ID,
		RoleID:           role.ID,
	})
	if err != nil {
		_, _ = s.organizationDirectory.DeleteOrganization(ctx, Organization.ID)
		return nil, err
	}
	return &Organization, nil
}

func (s *Service) GetOrganizationUserHasRole(ctx context.Context,
	userID int32) ([]access.Access, error) {
	return s.accessDirectory.GetAccessByUserID(ctx, userID)
}

func (s *Service) GetUserRoleInOrganization(ctx context.Context,
	userID int32, organizationName string) (int32, error) {
	Access, err := s.accessDirectory.GetAccessByOrganizationNameAndUserID(ctx,
		access.GetAccessByOrganizationNameAndUserIDParams{
			UserID:           userID,
			OrganizationName: organizationName,
		})
	if err != nil {
		return 0, err
	}
	return Access.RoleID, nil
}

func (s *Service) DeleteOrganization(ctx context.Context, id int32) error {
	_, err := s.organizationDirectory.DeleteOrganization(ctx, id)
	if err != nil {
		return err
	}
	_, err = s.accessDirectory.DeleteAccessByOrganizationID(ctx, id)
	return err
}
