package reservenow

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type ReserveNowReq struct {
	ConnectorID   int32            `json:"connectorId"`
	ExpiryDate    types.OcppTime `json:"expiryDate"`
	IDTag         string         `json:"idTag"`
	ParentIDTag   *string        `json:"parentIdTag,omitempty"`
	ReservationID int64          `json:"reservationID"`
}

type ReserveNowConf struct {
	Status db.ReservationStatus `json:"status"`
}

func unmarshalReserveNowConf(payload interface{}) (*ReserveNowConf, error) {
	reserveNowConf := &ReserveNowConf{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, reserveNowConf); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return reserveNowConf, nil
}
