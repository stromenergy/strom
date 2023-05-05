-- name: CreateConfiguration :one
INSERT INTO configurations (
    charge_point_id,
    key, 
    readonly,
    value,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: GetConfiguration :one
SELECT * FROM configurations
  WHERE id = $1;

-- name: GetConfigurationByKey :one
SELECT * FROM configurations
  WHERE charge_point_id = $1 AND key = $2;

-- name: ListConfigurations :many
SELECT * FROM configurations
  WHERE charge_point_id = $1
  ORDER BY id;

-- name: UpdateConfiguration :one
UPDATE configurations SET (
    value,
    updated_at
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
