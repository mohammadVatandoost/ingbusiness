package iam

import (
	"context"
	"fmt"
	"github.com/mohammadVatandoost/ingbusiness/internal/access"
	roles "github.com/mohammadVatandoost/ingbusiness/internal/role"
)

const (
	OwnerRole = iota + 1
	AdminRole
)

func (s *Service) GivePermissionByEmail(ctx context.Context,
	email string, organizationId int32) error {
	user, err := s.usersDirectory.GetUserByEmail(ctx, email)
	if err != nil {
		//ToDo: send email to user to register
		return fmt.Errorf("%s", ErrorUserIsNotRegistered)
	}

	//ToDo: get role type from user
	role, err := s.rolesDirectory.GetRoleByOrganizationIDAndRoleType(ctx,
		roles.GetRoleByOrganizationIDAndRoleTypeParams{
			OrganizationID: organizationId,
			RoleType:       AdminRole,
		})
	if err != nil {
		return fmt.Errorf("admin role does not exist, organizationId: %v", organizationId)
	}

	_, err = s.accessDirectory.AddAccess(ctx, access.AddAccessParams{
		UserID:         user.ID,
		OrganizationID: organizationId,
		RoleID:         role.ID,
	})
	return err
}

func (s *Service) DeletePermissionByEmail(ctx context.Context,
	email string, organizationId int32) error {
	user, err := s.usersDirectory.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	_, err = s.accessDirectory.DeleteAccessByOrganizationIDAndUserID(ctx,
		access.DeleteAccessByOrganizationIDAndUserIDParams{
			UserID:         user.ID,
			OrganizationID: organizationId,
		})
	return err
}
