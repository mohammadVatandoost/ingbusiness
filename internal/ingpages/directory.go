package ingpages

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

func (d *Directory) AddIngPage(ctx context.Context, arg AddIngPageParams) (IngPage, error) {
	ingPage, err := d.querier.AddIngPage(ctx, arg)
	if err != nil {
		return ingPage, status.Errorf(codes.Internal,
			"unexpected error AddIngPage: %s", err.Error())
	}
	return ingPage, nil
}

func (d *Directory) GetIngPage(ctx context.Context, id int32) (IngPage, error) {
	ingPage, err := d.querier.GetIngPage(ctx, id)
	if err != nil {
		return ingPage, status.Errorf(codes.Internal, "unexpected error GetIngPage: %s", err.Error())
	}
	return ingPage, nil
}

func (d *Directory) GetIngPages(ctx context.Context) ([]IngPage, error) {
	ingPages, err := d.querier.GetIngPages(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetIngPages: %s", err.Error())
	}
	return ingPages, nil
}

func (d *Directory) GetIngPageByUserID(ctx context.Context, userID int32) ([]IngPage, error) {
	ingPages, err := d.querier.GetIngPageByUserID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetIngPageByUserID: %s", err.Error())
	}
	return ingPages, nil
}

func (d *Directory) GetIngPageByOrganizationID(ctx context.Context, organizationID int32) ([]IngPage, error) {
	ingPages, err := d.querier.GetIngPageByOrganizationID(ctx, organizationID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetIngPageByOrganizationID: %s",
			err.Error())
	}
	return ingPages, nil
}

func (d *Directory) UpdateIngPageToken(ctx context.Context, arg UpdateIngPageTokenParams) (IngPage, error) {
	ingPage, err := d.querier.UpdateIngPageToken(ctx, arg)
	if err != nil {
		return ingPage, status.Errorf(codes.Internal, "unexpected error UpdateIngPageToken: %s", err.Error())
	}
	return ingPage, nil
}

func (d *Directory) DeleteIngPage(ctx context.Context, id int32) (IngPage, error) {
	ingPage, err := d.querier.DeleteIngPage(ctx, id)
	if err != nil {
		return ingPage, status.Errorf(codes.Internal, "unexpected error DeleteIngPage: %s", err.Error())
	}
	return ingPage, nil
}
