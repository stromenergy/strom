package types

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

type MessageCall struct {
	UniqueID string
	Action   Action
	Payload  interface{}
}

func NewMessageCall(uniqueID string, action Action, payload interface{}) *MessageCall {
	return &MessageCall{
		UniqueID: uniqueID,
		Action:   action,
		Payload:  payload,
	}
}

func (m *MessageCall) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{MessageTypeCALL, m.UniqueID, m.Action, m.Payload})
}

func (m *MessageCall) UnmarshalJSON(p []byte) error {
	var rawMessage []json.RawMessage
	var messageTypeID MessageType

	if err := json.Unmarshal(p, &rawMessage); err != nil {
		util.LogError("STR023: Error unmarshaling json", err)
		return err
	}

	if err := json.Unmarshal(rawMessage[0], &messageTypeID); err != nil {
		util.LogError("STR024: Error unmarshaling MessageTypeID", err)
		return err
	}

	if messageTypeID != MessageTypeCALL {
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

	m.Payload = ([]byte)(rawMessage[3])

	return nil
}

func (m *MessageCall) Send(client *ws.Client) {
	bytes, err := json.Marshal(m)

	if err != nil {
		util.LogError("STR028: Error marshaling call", err)
		return
	}

	client.Send(bytes)
}

type MessageCallResult struct {
	UniqueID string
	Payload  interface{}
}

func NewMessageCallResult(uniqueID string, payload interface{}) *MessageCallResult {
	return &MessageCallResult{
		UniqueID: uniqueID,
		Payload:  payload,
	}
}

func (m *MessageCallResult) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{MessageTypeCALLRESULT, m.UniqueID, m.Payload})
}

func (m *MessageCallResult) Send(client *ws.Client) {
	bytes, err := json.Marshal(m)

	if err != nil {
		util.LogError("STR029: Error marshaling result response", err)
		return
	}

	client.Send(bytes)
}

type NoError struct{}

type MessageCallError struct {
	UniqueID         string
	ErrorCode        ErrorCode
	ErrorDescription string
	ErrorDetails     interface{}
}

func NewMessageCallError(uniqueID string, code ErrorCode, description string, details interface{}) *MessageCallError {
	return &MessageCallError{
		UniqueID:         uniqueID,
		ErrorCode:        code,
		ErrorDescription: description,
		ErrorDetails:     details,
	}
}

func (m *MessageCallError) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{MessageTypeCALLERROR, m.UniqueID, m.ErrorCode, m.ErrorDescription, m.ErrorDetails})
}

func (m *MessageCallError) Send(client *ws.Client) {
	bytes, err := json.Marshal(m)

	if err != nil {
		util.LogError("STR030: Error marshaling error response", err)
		return
	}

	client.Send(bytes)
}
