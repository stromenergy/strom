-- name: CreateConnector :one
INSERT INTO connectors (
    connector_id,
    charge_point_id,
    error_code,
    status,
    info,
    vendor_id,
    vendor_error_code,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING *;

-- name: GetConnector :one
SELECT * FROM connectors
  WHERE id = $1;

-- name: GetConnectorByConnectorID :one
SELECT * FROM connectors
  WHERE charge_point_id = $1 AND connector_id = $2;

-- name: ListConnectors :many
SELECT * FROM connectors
  WHERE charge_point_id = $1
  ORDER BY connector_id;

-- name: UpdateConnector :one
UPDATE connectors SET (
    error_code,
    status,
    info,
    vendor_id,
    vendor_error_code,
    updated_at
  ) = ($2, $3, $4, $5, $6, $7)
  WHERE id = $1
  RETURNING *;
