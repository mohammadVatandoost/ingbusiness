package iam

import (
	"context"
	"fmt"
	"github.com/mohammadVatandoost/ingbusiness/internal/access"
)

const (
	OwnerRole = iota
	AdminRole
)

const ErrorUserIsNotRegistered = "کاربری با این ایمیل ثبت نام نشده است، از ایشان بخواهید در سایت ثبت نام کند."

func (s *Service) GivePermissionByEmail(ctx context.Context,
	email string, organizationId int32) error {
	user, err := s.usersDirectory.GetUserByEmail(ctx, email)
	if err != nil {
		//ToDo: send email to user to register
		return fmt.Errorf("%s", ErrorUserIsNotRegistered)
	}
	_, err = s.accessDirectory.AddAccess(ctx, access.AddAccessParams{
		UserID:         user.ID,
		OrganizationID: organizationId,
		RoleID:         AdminRole,
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
