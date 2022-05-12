-- name: AddSavedMessage :one
INSERT INTO saved_messages (
  message,
  ing_account_id,
  writer_id
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;


-- name: GetSavedMessages :many
SELECT * FROM saved_messages;

-- name: GetSavedMessage :one
SELECT * FROM saved_messages WHERE id = $1;

-- name: GetSavedMessageByWriterID :many
SELECT * FROM saved_messages WHERE writer_id = $1;

-- name: GetSavedMessageByIngAccountID :many
SELECT * FROM saved_messages WHERE ing_account_id = $1;

-- name: UpdateSavedMessageMessage :one
UPDATE saved_messages SET message = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSavedMessage :one
DELETE FROM saved_messages
WHERE id = $1
RETURNING *;

