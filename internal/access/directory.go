package access

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

func (d *Directory) AddAccess(ctx context.Context, arg AddAccessParams) (Access, error) {
	access, err := d.querier.AddAccess(ctx, arg)
	if err != nil {
		return access, status.Errorf(codes.Internal,
			"unexpected error AddAccess: %s", err.Error())
	}
	return access, nil
}

func (d *Directory) GetAccess(ctx context.Context, id int32) (Access, error) {
	access, err := d.querier.GetAccess(ctx, id)
	if err != nil {
		return access, status.Errorf(codes.Internal, "unexpected error GetAccess: %s", err.Error())
	}
	return access, nil
}

func (d *Directory) GetAccesses(ctx context.Context) ([]Access, error) {
	accesses, err := d.querier.GetAccesses(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetAccesses: %s", err.Error())
	}
	return accesses, nil
}

func (d *Directory) GetAccessByOrganizationID(ctx context.Context, ownerID int32) ([]Access, error) {
	accesses, err := d.querier.GetAccessByOrganizationID(ctx, ownerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetAccessByOrganizationID: %s",
			err.Error())
	}
	return accesses, nil
}

func (d *Directory) GetAccessByUserID(ctx context.Context, userID int32) ([]Access, error) {
	accesses, err := d.querier.GetAccessByUserID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetAccessByUserID: %s", err.Error())
	}
	return accesses, nil
}

func (d *Directory) GetAccessByOrganizationNameAndUserID(ctx context.Context, arg GetAccessByOrganizationNameAndUserIDParams) (Access, error) {
	access, err := d.querier.GetAccessByOrganizationNameAndUserID(ctx, arg)
	if err != nil {
		return access, status.Errorf(codes.Internal,
			"unexpected error GetAccessByOrganizationNameAndUserID: %s", err.Error())
	}
	return access, nil
}

func (d *Directory) DeleteAccess(ctx context.Context, id int32) (Access, error) {
	access, err := d.querier.DeleteAccess(ctx, id)
	if err != nil {
		return access, status.Errorf(codes.Internal, "unexpected error DeleteAccess: %s", err.Error())
	}
	return access, nil
}

func (d *Directory) DeleteAccessByOrganizationID(ctx context.Context, id int32) ([]Access, error) {
	accesses, err := d.querier.DeleteAccessByOrganizationID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error DeleteAccessByOrganizationID: %s",
			err.Error())
	}
	return accesses, nil
}

func (d *Directory) DeleteAccessByOrganizationIDAndUserID(ctx context.Context,
	arg DeleteAccessByOrganizationIDAndUserIDParams) ([]Access, error) {
	accesses, err := d.querier.DeleteAccessByOrganizationIDAndUserID(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error DeleteAccessByOrganizationIDAndUserID: %s",
			err.Error())
	}
	return accesses, nil
}

func (d *Directory) UpdateAccess(ctx context.Context, arg UpdateAccessParams) (Access, error) {
	access, err := d.querier.UpdateAccess(ctx, arg)
	if err != nil {
		return access, status.Errorf(codes.Internal, "unexpected error UpdateAccess: %s", err.Error())
	}
	return access, nil
}
