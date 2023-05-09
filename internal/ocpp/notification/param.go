package notification

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
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

func createConnectorParams(chargePointID int64, statusNotificationReq *StatusNotificationReq) db.CreateConnectorParams {
	timestamp := util.DefaultTime(types.NilTime(statusNotificationReq.Timestamp), time.Now())

	return db.CreateConnectorParams{
		ChargePointID:   chargePointID,
		ConnectorID:     statusNotificationReq.ConnectorID,
		ErrorCode:       db.ChargePointErrorCode(statusNotificationReq.ErrorCode),
		Status:          db.ConnectorStatus(statusNotificationReq.Status),
		Info:            util.SqlNullString(statusNotificationReq.Info),
		VendorID:        util.SqlNullString(statusNotificationReq.VendorID),
		VendorErrorCode: util.SqlNullString(statusNotificationReq.VendorErrorCode),
		CreatedAt:       timestamp,
		UpdatedAt:       timestamp,
	}
}
