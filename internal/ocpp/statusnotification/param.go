package statusnotification

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
)

func createConnectorParams(chargePointID int64, statusNotificationReq *StatusNotificationReq) db.CreateConnectorParams {
	timestamp := util.DefaultTime(types.NilTime(statusNotificationReq.Timestamp), time.Now())

	return db.CreateConnectorParams{
		ConnectorID:     statusNotificationReq.ConnectorID,
		ChargePointID:   chargePointID,
		ErrorCode:       db.ChargePointErrorCode(statusNotificationReq.ErrorCode),
		Status:          db.ChargePointStatus(statusNotificationReq.Status),
		Info:            util.SqlNullString(statusNotificationReq.Info),
		VendorID:        util.SqlNullString(statusNotificationReq.VendorID),
		VendorErrorCode: util.SqlNullString(statusNotificationReq.VendorErrorCode),
		CreatedAt:       timestamp,
		UpdatedAt:       timestamp,
	}
}
