package ocpp

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/util"
)

type MessageCall struct {
	MessageTypeID MessageType
	UniqueID      string
	Action        Action
	Payload       []byte
}

func (m *MessageCall) MarshalJSON() ([]byte, error) {
    return json.Marshal([]interface{}{m.MessageTypeID, m.UniqueID, m.Action, m.Payload})
}

func (m *MessageCall) UnmarshalJSON(p []byte) error {
    var rawMessage []json.RawMessage

    if err := json.Unmarshal(p, &rawMessage); err != nil {
        util.LogError("STR023: Error unmarshaling json", err)
        return err
    }

    if err := json.Unmarshal(rawMessage[0], &m.MessageTypeID); err != nil {
        util.LogError("STR024: Error unmarshaling MessageTypeID", err)
        return err
    }

    if m.MessageTypeID != MessageTypeCALL {
        err := errors.New("Message type mismatch")
        util.LogError("STR025: Message type mismatch", err)
        return err
    }

    if err := json.Unmarshal(rawMessage[1], &m.UniqueID); err != nil {
        util.LogError("STR026: Error unmarshaling UniqueID", err)
        return err
    }

    if err := json.Unmarshal(rawMessage[2], &m.Action); err != nil {
        util.LogError("STR027: Error unmarshaling Action", err)
        return err
    }

	m.Payload = rawMessage[3]
	
    return nil
}
