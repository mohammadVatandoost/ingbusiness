-- name: AddSavedMessage :one
INSERT INTO saved_messages (
  message,
  image,
  organization_id,
  writer_id
) VALUES (
  $1,
  $2,
  $3,
  $4
)
RETURNING *;


-- name: GetSavedMessages :many
SELECT * FROM saved_messages;

-- name: GetSavedMessage :one
SELECT * FROM saved_messages WHERE id = $1;

-- name: GetSavedMessageByWriterID :many
SELECT * FROM saved_messages WHERE writer_id = $1;

-- name: GetSavedMessageByOrganizationID :many
SELECT * FROM saved_messages WHERE organization_id = $1;

-- name: UpdateSavedMessageMessage :one
UPDATE saved_messages SET message = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSavedMessage :one
DELETE FROM saved_messages
WHERE id = $1
RETURNING *;

