package bootnotification

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/util"
)

func createChargePointParams(identity string, bootNotificationReq *BootNotificationReq) db.CreateChargePointParams {
	return db.CreateChargePointParams{
		Identity:          identity,
		Model:             bootNotificationReq.ChargePointModel,
		Vendor:            bootNotificationReq.ChargePointVendor,
		SerialNumber:      util.SqlNullString(bootNotificationReq.ChargePointSerialNumber),
		FirmwareVerion:    util.SqlNullString(bootNotificationReq.FirmwareVersion),
		ModemIccid:        util.SqlNullString(bootNotificationReq.Iccid),
		ModemImsi:         util.SqlNullString(bootNotificationReq.Imsi),
		MeterSerialNumber: util.SqlNullString(bootNotificationReq.MeterSerialNumber),
		MeterType:         util.SqlNullString(bootNotificationReq.MeterType),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}
