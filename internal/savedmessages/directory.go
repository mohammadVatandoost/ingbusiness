package savedmessages

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

func (d *Directory) AddSavedMessage(ctx context.Context, arg AddSavedMessageParams) (SavedMessage, error) {
	savedMessage, err := d.querier.AddSavedMessage(ctx, arg)
	if err != nil {
		return savedMessage, status.Errorf(codes.Internal,
			"unexpected error AddSavedMessage: %s", err.Error())
	}
	return savedMessage, nil
}

func (d *Directory) GetSavedMessage(ctx context.Context, id int32) (SavedMessage, error) {
	savedMessage, err := d.querier.GetSavedMessage(ctx, id)
	if err != nil {
		return savedMessage, status.Errorf(codes.Internal, "unexpected error GetSavedMessage: %s", err.Error())
	}
	return savedMessage, nil
}

func (d *Directory) GetSavedMessages(ctx context.Context) ([]SavedMessage, error) {
	savedMessages, err := d.querier.GetSavedMessages(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetSavedMessages: %s", err.Error())
	}
	return savedMessages, nil
}

func (d *Directory) GetSavedMessageByIngAccountID(ctx context.Context, ingAccountID int32) ([]SavedMessage, error) {
	savedMessages, err := d.querier.GetSavedMessageByIngAccountID(ctx, ingAccountID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetSavedMessageByIngAccountID: %s", err.Error())
	}
	return savedMessages, nil
}

func (d *Directory) GetSavedMessageByWriterID(ctx context.Context, writerID int32) ([]SavedMessage, error) {
	savedMessage, err := d.querier.GetSavedMessageByWriterID(ctx, writerID)
	if err != nil {
		return savedMessage, status.Errorf(codes.Internal, "unexpected error GetSavedMessageByWriterID: %s", err.Error())
	}
	return savedMessage, nil
}

func (d *Directory) DeleteIngAccount(ctx context.Context, id int32) (SavedMessage, error) {
	savedMessage, err := d.querier.DeleteSavedMessage(ctx, id)
	if err != nil {
		return savedMessage, status.Errorf(codes.Internal, "unexpected error DeleteSavedMessage: %s", err.Error())
	}
	return savedMessage, nil
}

func (d *Directory) UpdateSavedMessageMessage(ctx context.Context, arg UpdateSavedMessageMessageParams) (SavedMessage, error) {
	savedMessage, err := d.querier.UpdateSavedMessageMessage(ctx, arg)
	if err != nil {
		return savedMessage, status.Errorf(codes.Internal, "unexpected error UpdateSavedMessageMessage: %s", err.Error())
	}
	return savedMessage, nil
}
