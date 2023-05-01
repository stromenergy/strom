-- name: CreateReservation :one
INSERT INTO reservations (
    connector_id,
    charge_point_id,
    req_id,
    expiry_date,
    status,
    id_tag,
    parent_id_tag,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING *;

-- name: GetReservation :one
SELECT * FROM reservations
  WHERE id = $1;

-- name: GetReservationByReqID :one
SELECT * FROM reservations
  WHERE req_id = $1;

-- name: ListReservations :many
SELECT * FROM reservations
  ORDER BY id;

-- name: UpdateReservation :one
UPDATE reservations SET (
    req_id,
    status,
    updated_at
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
