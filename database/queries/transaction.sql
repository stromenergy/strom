-- name: CreateTransaction :one
INSERT INTO transactions (
    connector_id,
    charge_point_id,
    reservation_id,
    id_tag,
    meter_start,
    start_timestamp,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions
  WHERE id = $1;

-- name: ListTransactions :many
SELECT * FROM transactions
  ORDER BY id;

-- name: UpdateTransaction :one
UPDATE transactions SET (
    status,
    reason,
    meter_stop,
    stop_timestamp,
    updated_at
  ) = ($2, $3, $4, $5, $6)
  WHERE id = $1
  RETURNING *;
