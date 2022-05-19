package organization

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

func (d *Directory) AddOrganization(ctx context.Context, arg AddOrganizationParams) (Organization, error) {
	organization, err := d.querier.AddOrganization(ctx, arg)
	if err != nil {
		return organization, status.Errorf(codes.Internal,
			"unexpected error AddOrganization: %s", err.Error())
	}
	return organization, nil
}

func (d *Directory) GetOrganization(ctx context.Context, id int32) (Organization, error) {
	organization, err := d.querier.GetOrganization(ctx, id)
	if err != nil {
		return organization, status.Errorf(codes.Internal, "unexpected error GetOrganization: %s", err.Error())
	}
	return organization, nil
}

func (d *Directory) GetOrganizations(ctx context.Context) ([]Organization, error) {
	organizations, err := d.querier.GetOrganizations(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetOrganizations: %s", err.Error())
	}
	return organizations, nil
}

func (d *Directory) GetOrganizationByOwnerID(ctx context.Context, ownerID int32) ([]Organization, error) {
	organizations, err := d.querier.GetOrganizationByOwnerID(ctx, ownerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetOrganizationByOwnerID: %s", err.Error())
	}
	return organizations, nil
}

func (d *Directory) DeleteOrganization(ctx context.Context, id int32) (Organization, error) {
	savedMessage, err := d.querier.DeleteOrganization(ctx, id)
	if err != nil {
		return savedMessage, status.Errorf(codes.Internal, "unexpected error DeleteOrganization: %s", err.Error())
	}
	return savedMessage, nil
}

func (d *Directory) UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (Organization, error) {
	savedMessage, err := d.querier.UpdateOrganization(ctx, arg)
	if err != nil {
		return savedMessage, status.Errorf(codes.Internal, "unexpected error UpdateOrganization: %s", err.Error())
	}
	return savedMessage, nil
}
