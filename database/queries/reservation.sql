-- name: CreateReservation :one
INSERT INTO reservations (
    connector_id,
    charge_point_id,
    expiry_date,
    status,
    id_tag,
    parent_id_tag,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING *;

-- name: GetReservation :one
SELECT * FROM reservations
  WHERE id = $1;

-- name: ListReservations :many
SELECT * FROM reservations
  ORDER BY id;

-- name: UpdateReservation :one
UPDATE reservations SET (
    status,
    updated_at
  ) = ($2, $3)
  WHERE id = $1
  RETURNING *;
