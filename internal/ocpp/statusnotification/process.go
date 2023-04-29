package statusnotification

import (
	"context"
	"time"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/db/param"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *StatusNotification) StatusNotificationReq(client *ws.Client, message types.Message) {
	statusNotificationReq, err := unmarshalStatusNotificationReq(message.Payload)

	if err != nil {
		util.LogError("STR036: Error unmarshaling StatusNotificationReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	chargePoint, err := s.repository.GetChargePointByIdentity(ctx, client.ID)

	if err != nil {
		util.LogError("STR037: Charge point not found", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
		callError.Send(client)
		return
	}

	getConnectorByConnectorIDParams := db.GetConnectorByConnectorIDParams{
		ConnectorID:   statusNotificationReq.ConnectorID,
		ChargePointID: chargePoint.ID,
	}

	connector, err := s.repository.GetConnectorByConnectorID(ctx, getConnectorByConnectorIDParams)

	if err == nil {
		// Update existing connector
		updateConnectorParams := param.NewUpdateConnectorParams(connector)
		updateConnectorParams.ErrorCode = db.ChargePointErrorCode(statusNotificationReq.ErrorCode)
		updateConnectorParams.Status = db.ChargePointStatus(statusNotificationReq.Status)
		updateConnectorParams.Info = util.SqlNullString(statusNotificationReq.Info)
		updateConnectorParams.VendorID = util.SqlNullString(statusNotificationReq.VendorID)
		updateConnectorParams.VendorErrorCode = util.SqlNullString(statusNotificationReq.VendorErrorCode)
		updateConnectorParams.UpdatedAt = util.DefaultTime(types.NilTime(statusNotificationReq.Timestamp), time.Now())

		_, err = s.repository.UpdateConnector(ctx, updateConnectorParams)

		if err != nil {
			util.LogError("STR038: Error updating connector", err)
			callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
			callError.Send(client)
			return
		}
	} else {
		// Create new connector
		createConnectorParams := createConnectorParams(chargePoint.ID, statusNotificationReq)

		_, err = s.repository.CreateConnector(ctx, createConnectorParams)

		if err != nil {
			util.LogError("STR039: Error creating connector", err)
			callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeInternalError, "", types.NoError{})
			callError.Send(client)
			return
		}
	}

	callResult := types.NewMessageCallResult(message.UniqueID, StatusNotificationConf{})
	callResult.Send(client)

	// TODO: Notify UI of changes
}
