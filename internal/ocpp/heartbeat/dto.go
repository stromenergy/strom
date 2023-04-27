package heartbeat

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type HeartbeatConf struct {
	CurrentTime types.OcppTime `json:"currentTime"`
}
