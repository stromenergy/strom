package ocpp

import (
	"encoding/json"

	"github.com/stromenergy/strom/internal/util"
)

func (s *Ocpp) processMessage(message []byte) {
	messageCall := MessageCall{}

	if err := json.Unmarshal(message, &messageCall); err != nil {
        util.LogDebug("Message ignored")
		return
	}
}
