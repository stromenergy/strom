package management

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Management) SendChangeAvailabilityReq(client *ws.Client, connectorId int32, availabilityType types.AvailabilityType) (string, <-chan types.Message, error) {
	cancelReservationReq := ChangeAvailabilityReq{
		ConnectorID: connectorId,
		Type:        availabilityType,
	}

	return s.call.Send(client, types.CallActionChangeAvailability, cancelReservationReq)
}
