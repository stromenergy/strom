// Code generated by sqlc. DO NOT EDIT.
// source: charge_point.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createChargePoint = `-- name: CreateChargePoint :one
INSERT INTO charge_points (
    model, 
    vendor,
    serial_number,
    firmware_verion,
    modem_iccid,
    modem_imsi,
    meter_serial_number,
    meter_type,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING id, model, vendor, serial_number, firmware_verion, modem_iccid, modem_imsi, meter_serial_number, meter_type, created_at, updated_at
`

type CreateChargePointParams struct {
	Model             string         `db:"model" json:"model"`
	Vendor            string         `db:"vendor" json:"vendor"`
	SerialNumber      sql.NullString `db:"serial_number" json:"serialNumber"`
	FirmwareVerion    sql.NullString `db:"firmware_verion" json:"firmwareVerion"`
	ModemIccid        sql.NullString `db:"modem_iccid" json:"modemIccid"`
	ModemImsi         sql.NullString `db:"modem_imsi" json:"modemImsi"`
	MeterSerialNumber sql.NullString `db:"meter_serial_number" json:"meterSerialNumber"`
	MeterType         sql.NullString `db:"meter_type" json:"meterType"`
	CreatedAt         time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updatedAt"`
}

func (q *Queries) CreateChargePoint(ctx context.Context, arg CreateChargePointParams) (ChargePoint, error) {
	row := q.db.QueryRowContext(ctx, createChargePoint,
		arg.Model,
		arg.Vendor,
		arg.SerialNumber,
		arg.FirmwareVerion,
		arg.ModemIccid,
		arg.ModemImsi,
		arg.MeterSerialNumber,
		arg.MeterType,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i ChargePoint
	err := row.Scan(
		&i.ID,
		&i.Model,
		&i.Vendor,
		&i.SerialNumber,
		&i.FirmwareVerion,
		&i.ModemIccid,
		&i.ModemImsi,
		&i.MeterSerialNumber,
		&i.MeterType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getChargePoint = `-- name: GetChargePoint :one
SELECT id, model, vendor, serial_number, firmware_verion, modem_iccid, modem_imsi, meter_serial_number, meter_type, created_at, updated_at FROM charge_points
  WHERE id = $1
`

func (q *Queries) GetChargePoint(ctx context.Context, id int64) (ChargePoint, error) {
	row := q.db.QueryRowContext(ctx, getChargePoint, id)
	var i ChargePoint
	err := row.Scan(
		&i.ID,
		&i.Model,
		&i.Vendor,
		&i.SerialNumber,
		&i.FirmwareVerion,
		&i.ModemIccid,
		&i.ModemImsi,
		&i.MeterSerialNumber,
		&i.MeterType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listChargePoints = `-- name: ListChargePoints :many
SELECT id, model, vendor, serial_number, firmware_verion, modem_iccid, modem_imsi, meter_serial_number, meter_type, created_at, updated_at FROM charge_points
  ORDER BY id
`

func (q *Queries) ListChargePoints(ctx context.Context) ([]ChargePoint, error) {
	rows, err := q.db.QueryContext(ctx, listChargePoints)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ChargePoint
	for rows.Next() {
		var i ChargePoint
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Vendor,
			&i.SerialNumber,
			&i.FirmwareVerion,
			&i.ModemIccid,
			&i.ModemImsi,
			&i.MeterSerialNumber,
			&i.MeterType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateChargePoints = `-- name: UpdateChargePoints :one
UPDATE charge_points SET (
    model, 
    vendor,
    serial_number,
    firmware_verion,
    modem_iccid,
    modem_imsi,
    meter_serial_number,
    meter_type,
    updated_at
  ) = ($2, $3, $4, $5, $6, $7, $8, $9, $10)
  WHERE id = $1
  RETURNING id, model, vendor, serial_number, firmware_verion, modem_iccid, modem_imsi, meter_serial_number, meter_type, created_at, updated_at
`

type UpdateChargePointsParams struct {
	ID                int64          `db:"id" json:"id"`
	Model             string         `db:"model" json:"model"`
	Vendor            string         `db:"vendor" json:"vendor"`
	SerialNumber      sql.NullString `db:"serial_number" json:"serialNumber"`
	FirmwareVerion    sql.NullString `db:"firmware_verion" json:"firmwareVerion"`
	ModemIccid        sql.NullString `db:"modem_iccid" json:"modemIccid"`
	ModemImsi         sql.NullString `db:"modem_imsi" json:"modemImsi"`
	MeterSerialNumber sql.NullString `db:"meter_serial_number" json:"meterSerialNumber"`
	MeterType         sql.NullString `db:"meter_type" json:"meterType"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updatedAt"`
}

func (q *Queries) UpdateChargePoints(ctx context.Context, arg UpdateChargePointsParams) (ChargePoint, error) {
	row := q.db.QueryRowContext(ctx, updateChargePoints,
		arg.ID,
		arg.Model,
		arg.Vendor,
		arg.SerialNumber,
		arg.FirmwareVerion,
		arg.ModemIccid,
		arg.ModemImsi,
		arg.MeterSerialNumber,
		arg.MeterType,
		arg.UpdatedAt,
	)
	var i ChargePoint
	err := row.Scan(
		&i.ID,
		&i.Model,
		&i.Vendor,
		&i.SerialNumber,
		&i.FirmwareVerion,
		&i.ModemIccid,
		&i.ModemImsi,
		&i.MeterSerialNumber,
		&i.MeterType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
