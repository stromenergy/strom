-- name: CreateChargePoint :one
INSERT INTO charge_points (
    identity,
    model, 
    vendor,
    serial_number,
    firmware_verion,
    modem_iccid,
    modem_imsi,
    meter_serial_number,
    meter_type,
    status,
    password,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
  RETURNING *;

-- name: GetChargePoint :one
SELECT * FROM charge_points
  WHERE id = $1;

-- name: GetChargePointByIdentity :one
SELECT * FROM charge_points
  WHERE identity = $1;

-- name: ListChargePoints :many
SELECT * FROM charge_points
  ORDER BY id;

-- name: UpdateChargePoint :one
UPDATE charge_points SET (
    model, 
    vendor,
    serial_number,
    firmware_verion,
    modem_iccid,
    modem_imsi,
    meter_serial_number,
    meter_type,
    status,
    password,
    updated_at
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
  WHERE id = $1
  RETURNING *;
