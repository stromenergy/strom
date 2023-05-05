package management

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/ws"
)


func (s *Management) SendClearCacheReq(client *ws.Client) (string, <-chan types.Message) {
	return s.call.Send(client, types.CallActionClearCache, ClearCacheReq{})
}
