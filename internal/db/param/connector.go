package param

import (
	"time"

	"github.com/stromenergy/strom/internal/db"
)

func NewUpdateConnectorParams(connector db.Connector) db.UpdateConnectorParams {
	return db.UpdateConnectorParams{
		ID:              connector.ID,
		ErrorCode:       connector.ErrorCode,
		Status:          connector.Status,
		Info:            connector.Info,
		VendorID:        connector.VendorID,
		VendorErrorCode: connector.VendorErrorCode,
		UpdatedAt:       time.Now(),
	}
}
