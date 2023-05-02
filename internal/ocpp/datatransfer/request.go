package datatransfer

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *DataTransfer) Request(client *ws.Client, vendorId string, messageID, data *string) (string, <-chan types.Message) {
	dataTransferReq := DataTransferReq{
		VendorID:  vendorId,
		MessageID: messageID,
		Data:      data,
	}

	return s.call.Send(client, types.CallActionDataTransfer, dataTransferReq)
}
