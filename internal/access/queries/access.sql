-- name: AddAccess :one
INSERT INTO access (
    organization_id,
    organization_name,
    user_id,
    role_id
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;

-- name: GetAccesses :many
SELECT * FROM access;

-- name: GetAccess :one
SELECT * FROM access WHERE id = $1;

-- name: GetAccessByUserID :many
SELECT * FROM access WHERE user_id = $1;

-- name: GetAccessByOrganizationID :many
SELECT * FROM access WHERE organization_id = $1;

-- name: GetAccessByOrganizationNameAndUserID :one
SELECT * FROM access WHERE organization_name = $1 and user_id = $2;

-- name: UpdateAccess :one
UPDATE access SET role_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccess :one
DELETE FROM access
WHERE id = $1
RETURNING *;

-- name: DeleteAccessByOrganizationID :many
DELETE FROM access
WHERE organization_id = $1
RETURNING *;

-- name: DeleteAccessByOrganizationIDAndUserID :many
DELETE FROM access
WHERE organization_id = $1 and user_id = $2
RETURNING *;

