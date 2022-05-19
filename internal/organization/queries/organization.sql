-- name: AddOrganization :one
INSERT INTO organization (
    name,
    owner_id
) VALUES (
  $1,
  $2
)
RETURNING *;


-- name: GetOrganizations :many
SELECT * FROM organization;

-- name: GetOrganization :one
SELECT * FROM organization WHERE id = $1;

-- name: GetOrganizationByOwnerID :many
SELECT * FROM organization WHERE owner_id = $1;

-- name: UpdateOrganization :one
UPDATE organization SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrganization :one
DELETE FROM organization
WHERE id = $1
RETURNING *;

