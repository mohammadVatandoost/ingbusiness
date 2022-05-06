-- name: AddExperiment :one
INSERT INTO experiment (
  name,
  description,
  condition_id,
  condition_params,
  start_time,
  end_time,
  active
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
)
RETURNING *;

-- name: GetExperiments :many
SELECT * FROM experiment;

-- name: GetExperiment :one
SELECT * FROM experiment WHERE id = $1;

-- name: GetExperimentByName :one
SELECT * FROM experiment WHERE name = $1;

-- name: GetExperimentsByActiveState :many
SELECT * FROM experiment WHERE active = $1;

-- name: GetExperimentsExceeded :many
SELECT * FROM experiment WHERE active = true and end_time < $1;

-- name: GetExperimentsBySameConditions :many
SELECT * FROM experiment WHERE active = $1 and condition_id = $2 and condition_params = $3;

-- name: UpdateExperiment :one
UPDATE experiment SET name = $2, description = $3, condition_id = $4, condition_params = $5, start_time = $6, end_time = $7, active = $8
WHERE id = $1
RETURNING *;

-- name: DisableExperiment :one
UPDATE experiment SET active = false
WHERE id = $1
RETURNING *;

-- name: EnableExperiment :one
UPDATE experiment SET active = true
WHERE id = $1
RETURNING *;

-- name: DeleteExperiment :one
DELETE FROM experiment
WHERE id = $1
RETURNING *;

