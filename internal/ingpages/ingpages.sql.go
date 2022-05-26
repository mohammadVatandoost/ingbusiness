// Code generated by sqlc. DO NOT EDIT.
// source: ingpages.sql

package ingpages

import (
	"context"
)

const addIngPage = `-- name: AddIngPage :one
INSERT INTO ing_pages (
  name,
  token,
  organization_id,
  creator_id
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING id, name, token, organization_id, creator_id, create_time, update_time
`

type AddIngPageParams struct {
	Name           string
	Token          string
	OrganizationID int32
	CreatorID      int32
}

func (q *Queries) AddIngPage(ctx context.Context, arg AddIngPageParams) (IngPage, error) {
	row := q.db.QueryRowContext(ctx, addIngPage,
		arg.Name,
		arg.Token,
		arg.OrganizationID,
		arg.CreatorID,
	)
	var i IngPage
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.OrganizationID,
		&i.CreatorID,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const deleteIngPage = `-- name: DeleteIngPage :one
DELETE FROM ing_pages
WHERE id = $1
RETURNING id, name, token, organization_id, creator_id, create_time, update_time
`

func (q *Queries) DeleteIngPage(ctx context.Context, id int32) (IngPage, error) {
	row := q.db.QueryRowContext(ctx, deleteIngPage, id)
	var i IngPage
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.OrganizationID,
		&i.CreatorID,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getIngPage = `-- name: GetIngPage :one
SELECT id, name, token, organization_id, creator_id, create_time, update_time FROM ing_pages WHERE id = $1
`

func (q *Queries) GetIngPage(ctx context.Context, id int32) (IngPage, error) {
	row := q.db.QueryRowContext(ctx, getIngPage, id)
	var i IngPage
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.OrganizationID,
		&i.CreatorID,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}

const getIngPageByOrganizationID = `-- name: GetIngPageByOrganizationID :many
SELECT id, name, token, organization_id, creator_id, create_time, update_time FROM ing_pages WHERE organization_id = $1
`

func (q *Queries) GetIngPageByOrganizationID(ctx context.Context, organizationID int32) ([]IngPage, error) {
	rows, err := q.db.QueryContext(ctx, getIngPageByOrganizationID, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IngPage
	for rows.Next() {
		var i IngPage
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Token,
			&i.OrganizationID,
			&i.CreatorID,
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

const getIngPageByUserID = `-- name: GetIngPageByUserID :many
SELECT id, name, token, organization_id, creator_id, create_time, update_time FROM ing_pages WHERE creator_id = $1
`

func (q *Queries) GetIngPageByUserID(ctx context.Context, creatorID int32) ([]IngPage, error) {
	rows, err := q.db.QueryContext(ctx, getIngPageByUserID, creatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IngPage
	for rows.Next() {
		var i IngPage
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Token,
			&i.OrganizationID,
			&i.CreatorID,
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

const getIngPages = `-- name: GetIngPages :many
SELECT id, name, token, organization_id, creator_id, create_time, update_time FROM ing_pages
`

func (q *Queries) GetIngPages(ctx context.Context) ([]IngPage, error) {
	rows, err := q.db.QueryContext(ctx, getIngPages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []IngPage
	for rows.Next() {
		var i IngPage
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Token,
			&i.OrganizationID,
			&i.CreatorID,
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

const updateIngPageToken = `-- name: UpdateIngPageToken :one
UPDATE ing_pages SET token = $2
WHERE id = $1
RETURNING id, name, token, organization_id, creator_id, create_time, update_time
`

type UpdateIngPageTokenParams struct {
	ID    int32
	Token string
}

func (q *Queries) UpdateIngPageToken(ctx context.Context, arg UpdateIngPageTokenParams) (IngPage, error) {
	row := q.db.QueryRowContext(ctx, updateIngPageToken, arg.ID, arg.Token)
	var i IngPage
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Token,
		&i.OrganizationID,
		&i.CreatorID,
		&i.CreateTime,
		&i.UpdateTime,
	)
	return i, err
}