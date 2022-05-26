-- name: AddIngPage :one
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
RETURNING *;

-- name: GetIngPages :many
SELECT * FROM ing_pages;

-- name: GetIngPage :one
SELECT * FROM ing_pages WHERE id = $1;

-- name: GetIngPageByUserID :many
SELECT * FROM ing_pages WHERE creator_id = $1;

-- name: GetIngPageByOrganizationID :many
SELECT * FROM ing_pages WHERE organization_id = $1;

-- name: UpdateIngPageToken :one
UPDATE ing_pages SET token = $2
WHERE id = $1
RETURNING *;

-- name: DeleteIngPage :one
DELETE FROM ing_pages
WHERE id = $1
RETURNING *;

