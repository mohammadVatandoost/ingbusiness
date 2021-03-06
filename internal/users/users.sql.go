// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package users

import (
	"context"
)

const addUser = `-- name: AddUser :one
INSERT INTO users (
  name,
  username,
  email,
  phone,
  password,
  profile_image
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING id, name, username, email, phone, password, profile_image, create_time, update_time
`

type AddUserParams struct {
	Name         string
	Username     string
	Email        string
	Phone        string
	Password     string
	ProfileImage string
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, addUser,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Phone,
		arg.Password,
		arg.ProfileImage,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.ProfileImage,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id, name, username, email, phone, password, profile_image, create_time, update_time
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.ProfileImage,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name, username, email, phone, password, profile_image, create_time, update_time FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.ProfileImage,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, username, email, phone, password, profile_image, create_time, update_time FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.ProfileImage,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, name, username, email, phone, password, profile_image, create_time, update_time FROM users WHERE phone = $1
`

func (q *Queries) GetUserByPhone(ctx context.Context, phone string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhone, phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.ProfileImage,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getUserByUserName = `-- name: GetUserByUserName :one
SELECT id, name, username, email, phone, password, profile_image, create_time, update_time FROM users WHERE username = $1
`

func (q *Queries) GetUserByUserName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUserName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.ProfileImage,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, name, username, email, phone, password, profile_image, create_time, update_time FROM users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Email,
			&i.Phone,
			&i.Password,
			&i.ProfileImage,
			&i.CreateTime,
			&i.UpdateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET name = $2, username = $3, email = $4, phone = $5, password = $6
WHERE id = $1
RETURNING id, name, username, email, phone, password, profile_image, create_time, update_time
`

type UpdateUserParams struct {
	ID       int32
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Username,
		arg.Email,
		arg.Phone,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.ProfileImage,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}
