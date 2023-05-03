package ocpp

import (
	"time"

	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Ocpp) ChangeAvailability(client *ws.Client, connectorId int32, availabilityType types.AvailabilityType) (string, <-chan types.Message) {
	return s.notification.SendChangeAvailabilityReq(client, connectorId, availabilityType)
}

func (s *Ocpp) CancelReservation(client *ws.Client, reservationID int64) {
	s.reservation.SendCancelReservationReq(client, reservationID)
}

func (s *Ocpp) DataTransfer(client *ws.Client, vendorId string, messageID, data *string) (string, <-chan types.Message) {
	return s.dataTransfer.SendDataTransferReq(client, vendorId, messageID, data)
}

func (s *Ocpp) ReserveNow(client *ws.Client, connectorId int32, expiryDate time.Time, idTag string, parentIdTag *string) {
	s.reservation.SendReserveNowReq(client, connectorId, expiryDate, idTag, parentIdTag)
}

func (s *Ocpp) TriggerMessage(client *ws.Client, chargePointId int64, messageTrigger types.MessageTrigger, connectorId *int32) (string, <-chan types.Message) {
	return s.triggerMessage.SendTriggerMessageReq(client, chargePointId, messageTrigger, connectorId)
}