package ocpp

func (s *Ocpp) bootNotificationReq(messageCall MessageCall) (*MessageCallResult, *MessageCallError) {
	return nil, NewMessageCallError(messageCall.UniqueID, ErrorCodeNOTIMPLEMENTED, "", NoError{})
}