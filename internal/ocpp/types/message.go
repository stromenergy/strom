package types

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

type NoError struct{}

type Message struct {
	MessageType      MessageType
	UniqueID         string
	Action           CallAction
	ErrorCode        ErrorCode
	ErrorDescription string
	Payload          interface{}
}

func NewMessageCall(uniqueID string, action CallAction, payload interface{}) *Message {
	return &Message{
		MessageType: MessageTypeCall,
		UniqueID:    uniqueID,
		Action:      action,
		Payload:     payload,
	}
}

func NewMessageCallError(uniqueID string, code ErrorCode, description string, details interface{}) *Message {
	if details == nil {
		details = NoError{}
	}

	return &Message{
		MessageType:      MessageTypeCallError,
		UniqueID:         uniqueID,
		ErrorCode:        code,
		ErrorDescription: description,
		Payload:          details,
	}
}

func NewMessageCallResult(uniqueID string, payload interface{}) *Message {
	return &Message{
		MessageType: MessageTypeCallResult,
		UniqueID:    uniqueID,
		Payload:     payload,
	}
}

func (m *Message) MarshalJSON() ([]byte, error) {
	switch m.MessageType {
	case MessageTypeCall:
		return json.Marshal([]interface{}{MessageTypeCall, m.UniqueID, m.Action, m.Payload})
	case MessageTypeCallError:
		return json.Marshal([]interface{}{MessageTypeCall, m.UniqueID, m.ErrorCode, m.ErrorDescription, m.Payload})
	case MessageTypeCallResult:
		return json.Marshal([]interface{}{MessageTypeCall, m.UniqueID, m.Payload})
	}

	return nil, errors.New("Invalid message type")
}

func (m *Message) UnmarshalJSON(p []byte) error {
	var rawMessage []json.RawMessage

	if err := json.Unmarshal(p, &rawMessage); err != nil {
		util.LogError("STR023: Error unmarshaling json", err)
		return err
	}

	if err := json.Unmarshal(rawMessage[0], &m.MessageType); err != nil {
		util.LogError("STR024: Error unmarshaling MessageTypeID", err)
		return err
	}

	if err := json.Unmarshal(rawMessage[1], &m.UniqueID); err != nil {
		util.LogError("STR025: Error unmarshaling UniqueID", err)
		return err
	}

	switch m.MessageType {
	case MessageTypeCall:
		if err := json.Unmarshal(rawMessage[2], &m.Action); err != nil {
			util.LogError("STR026: Error unmarshaling Action", err)
			return err
		}

		m.Payload = ([]byte)(rawMessage[3])
	case MessageTypeCallError:
		if err := json.Unmarshal(rawMessage[2], &m.ErrorCode); err != nil {
			util.LogError("STR027: Error unmarshaling ErrorCode", err)
			return err
		}

		if err := json.Unmarshal(rawMessage[3], &m.ErrorDescription); err != nil {
			util.LogError("STR028: Error unmarshaling ErrorDescription", err)
			return err
		}

		m.Payload = ([]byte)(rawMessage[4])
	case MessageTypeCallResult:
		m.Payload = ([]byte)(rawMessage[2])
	}

	return nil
}

func (m *Message) Send(client *ws.Client) error {
	bytes, err := json.Marshal(m)

	if err != nil {
		util.LogError("STR029: Error marshaling message", err)
		return errors.New("Error sending message")
	}

	client.Send(bytes)

	return nil
}
