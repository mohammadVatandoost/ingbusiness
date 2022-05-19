package users

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

func (d *Directory) AddUser(ctx context.Context, arg AddUserParams) (User, error) {
	user, err := d.querier.AddUser(ctx, arg)
	if err != nil {
		return user, status.Errorf(codes.Internal,
			"unexpected error AddUser: %s", err.Error())
	}
	return user, nil
}

func (d *Directory) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	user, err := d.querier.UpdateUser(ctx, arg)
	if err != nil {
		return user, status.Errorf(codes.Internal, "unexpected error UpdateUser: %s", err.Error())
	}
	return user, nil
}

func (d *Directory) GetUsers(ctx context.Context) ([]User, error) {
	users, err := d.querier.GetUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unexpected error GetUsers: %s", err.Error())
	}
	return users, nil
}

func (d *Directory) GetUser(ctx context.Context, id int32) (User, error) {
	user, err := d.querier.GetUser(ctx, id)
	if err != nil {
		return user, status.Errorf(codes.Internal, "unexpected error GetUser: %s", err.Error())
	}
	return user, nil
}

func (d *Directory) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user, err := d.querier.GetUserByEmail(ctx, email)
	if err != nil {
		return user, status.Errorf(codes.Internal, "unexpected error GetUserByEmail: %s", err.Error())
	}
	return user, nil
}

func (d *Directory) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	user, err := d.querier.GetUserByPhone(ctx, phone)
	if err != nil {
		return user, status.Errorf(codes.Internal, "unexpected error GetUserByPhone: %s", err.Error())
	}
	return user, nil
}

func (d *Directory) GetUserByUserName(ctx context.Context,
	username string) (User, error) {
	user, err := d.querier.GetUserByUserName(ctx, username)
	if err != nil {
		return user,
			status.Errorf(codes.Internal, "unexpected error GetUserByUserName: %s",
				err.Error())
	}
	return user, nil
}

func (d *Directory) DeleteUser(ctx context.Context, id int32) (User, error) {
	user, err := d.querier.DeleteUser(ctx, id)
	if err != nil {
		return user, status.Errorf(codes.Internal, "unexpected error DeleteUser: %s", err.Error())
	}
	return user, nil
}
