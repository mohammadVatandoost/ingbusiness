-- name: AddIngAccount :one
INSERT INTO ing_accounts (
  name,
  token,
  owner_id
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetIngAccounts :many
SELECT * FROM ing_accounts;

-- name: GetIngAccount :one
SELECT * FROM ing_accounts WHERE id = $1;

-- name: GetIngAccountByUserID :one
SELECT * FROM ing_accounts WHERE owner_id = $1;

-- name: UpdateIngAccountToken :one
UPDATE ing_accounts SET token = $2
WHERE id = $1
RETURNING *;

-- name: DeleteIngAccount :one
DELETE FROM ing_accounts
WHERE id = $1
RETURNING *;

