package roles

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Directory struct {
	logger  *logrus.Logger
	db      *sql.DB
	sb      squirrel.StatementBuilderType
	querier Querier
}

func NewDirectory(logger *logrus.Logger, db *sql.DB) *Directory {
	return &Directory{
		logger:  logger,
		db:      db,
		sb:      squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
		querier: New(db),
	}
}

func (d *Directory) Close() error {
	return d.db.Close()
}

func (d *Directory) AddRole(ctx context.Context, arg AddRoleParams) (Role, error) {
	role, err := d.querier.AddRole(ctx, arg)
	if err != nil {
		return role, status.Errorf(codes.Internal,
			"unexpected error AddRole: %s", err.Error())
	}
	return role, nil
}

func (d *Directory) GetRole(ctx context.Context, id int32) (Role, error) {
	role, err := d.querier.GetRole(ctx, id)
	if err != nil {
		return role, status.Errorf(codes.Internal, "unexpected error GetRole: %s", err.Error())
	}
	return role, nil
}

func (d *Directory) GetRoles(ctx context.Context) ([]Role, error) {
	roles, err := d.querier.GetRoles(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetRoles: %s", err.Error())
	}
	return roles, nil
}

func (d *Directory) GetRoleByOrganizationID(ctx context.Context, ownerID int32) ([]Role, error) {
	roles, err := d.querier.GetRoleByOrganizationID(ctx, ownerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetRoleByOrganizationID: %s", err.Error())
	}
	return roles, nil
}

func (d *Directory) GetRoleByCreatorID(ctx context.Context, creatorID int32) ([]Role, error) {
	roles, err := d.querier.GetRoleByCreatorID(ctx, creatorID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetRoleByCreatorID: %s", err.Error())
	}
	return roles, nil
}

func (d *Directory) DeleteRole(ctx context.Context, id int32) (Role, error) {
	role, err := d.querier.DeleteRole(ctx, id)
	if err != nil {
		return role, status.Errorf(codes.Internal, "unexpected error DeleteRole: %s", err.Error())
	}
	return role, nil
}
