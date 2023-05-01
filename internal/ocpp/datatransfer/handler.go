package datatransfer

import (
	"fmt"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
)

type Handler interface {
	DataTransferReq(chargePoint db.ChargePoint, dataTransferReq *DataTransferReq) (types.DataTransferStatus, *string)
}

func (s *DataTransfer) AddHandler(vendorId string, messageId *string, handler Handler) {
	key := getHandlerKey(vendorId, messageId)
	s.handlers[key] = handler
}

func getHandlerKey(vendorId string, messageId *string) string {
	return fmt.Sprintf("%s:%s", vendorId, util.DefaultString(messageId, "*"))
}
