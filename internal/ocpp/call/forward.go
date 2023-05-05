package call

import (
	"github.com/stromenergy/strom/internal/ocpp/types"
)

func Forward(message types.Message, channel chan<- types.Message) {
	channel <- message
}
