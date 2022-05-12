package ingaccounts

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

func (d *Directory) AddIngAccount(ctx context.Context, arg AddIngAccountParams) (IngAccount, error) {
	ingAccount, err := d.querier.AddIngAccount(ctx, arg)
	if err != nil {
		return ingAccount, status.Errorf(codes.Internal,
			"unexpected error AddIngAccount: %s", err.Error())
	}
	return ingAccount, nil
}

func (d *Directory) GetIngAccount(ctx context.Context, id int32) (IngAccount, error) {
	ingAccount, err := d.querier.GetIngAccount(ctx, id)
	if err != nil {
		return ingAccount, status.Errorf(codes.Internal, "unexpected error GetIngAccount: %s", err.Error())
	}
	return ingAccount, nil
}

func (d *Directory) GetIngAccounts(ctx context.Context) ([]IngAccount, error) {
	ingAccounts, err := d.querier.GetIngAccounts(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetIngAccounts: %s", err.Error())
	}
	return ingAccounts, nil
}

func (d *Directory) GetIngAccountByUserID(ctx context.Context, userID int32) (IngAccount, error) {
	ingAccount, err := d.querier.GetIngAccountByUserID(ctx, userID)
	if err != nil {
		return ingAccount, status.Errorf(codes.Internal, "unexpected error GetIngAccountByUserID: %s", err.Error())
	}
	return ingAccount, nil
}

func (d *Directory) UpdateIngAccountToken(ctx context.Context, arg UpdateIngAccountTokenParams) (IngAccount, error) {
	ingAccount, err := d.querier.UpdateIngAccountToken(ctx, arg)
	if err != nil {
		return ingAccount, status.Errorf(codes.Internal, "unexpected error UpdateIngAccountToken: %s", err.Error())
	}
	return ingAccount, nil
}

func (d *Directory) DeleteIngAccount(ctx context.Context, id int32) (IngAccount, error) {
	ingAccount, err := d.querier.DeleteIngAccount(ctx, id)
	if err != nil {
		return ingAccount, status.Errorf(codes.Internal, "unexpected error DeleteIngAccount: %s", err.Error())
	}
	return ingAccount, nil
}
