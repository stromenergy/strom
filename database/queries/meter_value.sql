-- name: CreateMeterValue :one
INSERT INTO meter_values (
    connector_id,
    charge_point_id,
    transaction_id,
    format,
    context,
    measurand,
    phase,
    location,
    unit,
    raw_value,
    signed_data_value,
    timestamp,
    created_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  RETURNING *;

-- name: GetMeterValue :one
SELECT * FROM meter_values
  WHERE id = $1;

-- name: ListMeterValues :many
SELECT * FROM meter_values
  ORDER BY id;
