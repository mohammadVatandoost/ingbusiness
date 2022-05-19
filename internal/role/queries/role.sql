-- name: AddRole :one
INSERT INTO roles (
    organization_id,
    creator_id,
    role_type
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetRoles :many
SELECT * FROM roles;

-- name: GetRole :one
SELECT * FROM roles WHERE id = $1;

-- name: GetRoleByOrganizationID :many
SELECT * FROM roles WHERE organization_id = $1;

-- name: GetRoleByCreatorID :many
SELECT * FROM roles WHERE creator_id = $1;

-- name: DeleteRole :one
DELETE FROM roles
WHERE id = $1
RETURNING *;

