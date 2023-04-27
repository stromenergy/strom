package ocpp

import (
)

type BootNotificationConf struct {
	CurrentTime OcppTime          `json:"currentTime"`
	Interval    int                `json:"interval"`
	Status      RegistrationStatus `json:"status"`
}

func (s *Ocpp) bootNotificationReq(messageCall MessageCall) (*MessageCallResult, *MessageCallError) {
	// TODO unmarshal data and update db

	bootNotificationConf := BootNotificationConf{
		CurrentTime: NewOcppTime(nil),
		Interval: 900,
		Status: RegistrationStatusACCEPTED,
	}

	return NewMessageCallResult(messageCall.UniqueID, bootNotificationConf), nil
}
