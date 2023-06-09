// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: meter_value.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMeterValue = `-- name: CreateMeterValue :one
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
  RETURNING id, connector_id, charge_point_id, transaction_id, format, context, measurand, phase, location, unit, raw_value, signed_data_value, timestamp, created_at
`

type CreateMeterValueParams struct {
	ConnectorID     int32                  `db:"connector_id" json:"connectorID"`
	ChargePointID   int64                  `db:"charge_point_id" json:"chargePointID"`
	TransactionID   sql.NullInt64          `db:"transaction_id" json:"transactionID"`
	Format          MeterValueFormat       `db:"format" json:"format"`
	Context         MeterReadingContext    `db:"context" json:"context"`
	Measurand       MeterMeasurand         `db:"measurand" json:"measurand"`
	Phase           NullMeterPhase         `db:"phase" json:"phase"`
	Location        MeterLocation          `db:"location" json:"location"`
	Unit            NullMeterUnitOfMeasure `db:"unit" json:"unit"`
	RawValue        sql.NullFloat64        `db:"raw_value" json:"rawValue"`
	SignedDataValue sql.NullString         `db:"signed_data_value" json:"signedDataValue"`
	Timestamp       time.Time              `db:"timestamp" json:"timestamp"`
	CreatedAt       time.Time              `db:"created_at" json:"createdAt"`
}

func (q *Queries) CreateMeterValue(ctx context.Context, arg CreateMeterValueParams) (MeterValue, error) {
	row := q.db.QueryRowContext(ctx, createMeterValue,
		arg.ConnectorID,
		arg.ChargePointID,
		arg.TransactionID,
		arg.Format,
		arg.Context,
		arg.Measurand,
		arg.Phase,
		arg.Location,
		arg.Unit,
		arg.RawValue,
		arg.SignedDataValue,
		arg.Timestamp,
		arg.CreatedAt,
	)
	var i MeterValue
	err := row.Scan(
		&i.ID,
		&i.ConnectorID,
		&i.ChargePointID,
		&i.TransactionID,
		&i.Format,
		&i.Context,
		&i.Measurand,
		&i.Phase,
		&i.Location,
		&i.Unit,
		&i.RawValue,
		&i.SignedDataValue,
		&i.Timestamp,
		&i.CreatedAt,
	)
	return i, err
}

const getMeterValue = `-- name: GetMeterValue :one
SELECT id, connector_id, charge_point_id, transaction_id, format, context, measurand, phase, location, unit, raw_value, signed_data_value, timestamp, created_at FROM meter_values
  WHERE id = $1
`

func (q *Queries) GetMeterValue(ctx context.Context, id int64) (MeterValue, error) {
	row := q.db.QueryRowContext(ctx, getMeterValue, id)
	var i MeterValue
	err := row.Scan(
		&i.ID,
		&i.ConnectorID,
		&i.ChargePointID,
		&i.TransactionID,
		&i.Format,
		&i.Context,
		&i.Measurand,
		&i.Phase,
		&i.Location,
		&i.Unit,
		&i.RawValue,
		&i.SignedDataValue,
		&i.Timestamp,
		&i.CreatedAt,
	)
	return i, err
}

const listMeterValues = `-- name: ListMeterValues :many
SELECT id, connector_id, charge_point_id, transaction_id, format, context, measurand, phase, location, unit, raw_value, signed_data_value, timestamp, created_at FROM meter_values
  ORDER BY id
`

func (q *Queries) ListMeterValues(ctx context.Context) ([]MeterValue, error) {
	rows, err := q.db.QueryContext(ctx, listMeterValues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MeterValue
	for rows.Next() {
		var i MeterValue
		if err := rows.Scan(
			&i.ID,
			&i.ConnectorID,
			&i.ChargePointID,
			&i.TransactionID,
			&i.Format,
			&i.Context,
			&i.Measurand,
			&i.Phase,
			&i.Location,
			&i.Unit,
			&i.RawValue,
			&i.SignedDataValue,
			&i.Timestamp,
			&i.CreatedAt,
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
