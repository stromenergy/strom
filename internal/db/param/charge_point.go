package param

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
)

func NewUpdateChargePointParams(chargePoint db.ChargePoint) db.UpdateChargePointParams {
	return db.UpdateChargePointParams{
		ID:                chargePoint.ID,
		Model:             chargePoint.Model,
		Vendor:            chargePoint.Vendor,
		SerialNumber:      chargePoint.SerialNumber,
		FirmwareVerion:    chargePoint.FirmwareVerion,
		ModemIccid:        chargePoint.ModemIccid,
		ModemImsi:         chargePoint.ModemImsi,
		MeterSerialNumber: chargePoint.MeterSerialNumber,
		MeterType:         chargePoint.MeterType,
		Status:            chargePoint.Status,
		Password:          chargePoint.Password,
		UpdatedAt:         time.Now(),
	}
}
