package ocpp

import (
	"time"

	"github.com/stromenergy/strom/internal/ocpp/transaction"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) ChangeAvailability(client *ws.Client, connectorId int32, availabilityType types.AvailabilityType) (string, <-chan types.Message) {
	return s.management.SendChangeAvailabilityReq(client, connectorId, availabilityType)
}

func (s *Ocpp) ChangeConfiguration(client *ws.Client, key, value string) {
	s.management.SendChangeConfigurationReq(client, key, value)
}

func (s *Ocpp) ClearCache(client *ws.Client) (string, <-chan types.Message) {
	return s.management.SendClearCacheReq(client)
}

func (s *Ocpp) CancelReservation(client *ws.Client, reservationID int64) {
	s.reservation.SendCancelReservationReq(client, reservationID)
}

func (s *Ocpp) DataTransfer(client *ws.Client, vendorId string, messageID, data *string) (string, <-chan types.Message) {
	return s.dataTransfer.SendDataTransferReq(client, vendorId, messageID, data)
}

func (s *Ocpp) GetConfiguration(client *ws.Client, keys *[]string) {
	s.management.SendGetConfigurationReq(client, keys)
}

func (s *Ocpp) RemoteStartTransaction(client *ws.Client, connectorId *int32, idTag string, chargingProfile *transaction.ChargingProfile) (string, <-chan types.Message) {
	return s.transaction.SendRemoteStartTransactionReq(client, connectorId, idTag, chargingProfile)
}

func (s *Ocpp) RemoteStopTransaction(client *ws.Client, transactionID int64) (string, <-chan types.Message) {
	return s.transaction.SendRemoteStopTransactionReq(client, transactionID)
}

func (s *Ocpp) ReserveNow(client *ws.Client, connectorId int32, expiryDate time.Time, idTag string, parentIdTag *string) {
	s.reservation.SendReserveNowReq(client, connectorId, expiryDate, idTag, parentIdTag)
}

func (s *Ocpp) Reset(client *ws.Client, resetType types.ResetType) (string, <-chan types.Message) {
	return s.management.SendResetReq(client, resetType)
}

func (s *Ocpp) TriggerMessage(client *ws.Client, chargePointId int64, messageTrigger types.MessageTrigger, connectorId *int32) (string, <-chan types.Message) {
	return s.triggerMessage.SendTriggerMessageReq(client, chargePointId, messageTrigger, connectorId)
}

func (s *Ocpp) UnlockConnector(client *ws.Client, connectorId int32) (string, <-chan types.Message) {
	return s.management.SendUnlockConnectorReq(client, connectorId)
}
