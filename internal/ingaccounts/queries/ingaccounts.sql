-- name: AddIngAccount :one
INSERT INTO ing_accounts (
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

-- name: GetIngAccounts :many
SELECT * FROM ing_accounts;

-- name: GetIngAccount :one
SELECT * FROM ing_accounts WHERE id = $1;

-- name: GetIngAccountByUserID :one
SELECT * FROM ing_accounts WHERE creator_id = $1;

-- name: GetIngAccountByOrganizationID :one
SELECT * FROM ing_accounts WHERE organization_id = $1;

-- name: UpdateIngAccountToken :one
UPDATE ing_accounts SET token = $2
WHERE id = $1
RETURNING *;

-- name: DeleteIngAccount :one
DELETE FROM ing_accounts
WHERE id = $1
RETURNING *;

