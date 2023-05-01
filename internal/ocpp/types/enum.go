package types

type DataTransferStatus string

const (
	DataTransferStatusAccepted         DataTransferStatus = "Accepted"
	DataTransferStatusRejected         DataTransferStatus = "Rejected"
	DataTransferStatusUnknownMessageId DataTransferStatus = "UnknownMessageId"
	DataTransferStatusUnknownVendorId  DataTransferStatus = "UnknownVendorId"
)

type ErrorCode string

const (
	ErrorCodeNotImplemented               ErrorCode = "NotImplemented"
	ErrorCodeNotSupported                 ErrorCode = "NotSupported"
	ErrorCodeInternalError                ErrorCode = "InternalError"
	ErrorCodeProtocolError                ErrorCode = "ProtocolError"
	ErrorCodeSecurityError                ErrorCode = "SecurityError"
	ErrorCodeFormationViolation           ErrorCode = "FormationViolation"
	ErrorCodePropertyConstraintViolation  ErrorCode = "PropertyConstraintViolation"
	ErrorCodeOccurenceConstraintViolation ErrorCode = "OccurenceConstraintViolation"
	ErrorCodeTypeConstraintViolation      ErrorCode = "TypeConstraintViolation"
	ErrorCodeGenericError                 ErrorCode = "GenericError"
)

type MessageType int

const (
	MessageTypeCall       MessageType = 2
	MessageTypeCallResult MessageType = 3
	MessageTypeCallError  MessageType = 4
)

type MessageTrigger string

const (
	MessageTriggerBootNotification              MessageTrigger = "BootNotification"
	MessageTriggerDiagnosticsStatusNotification MessageTrigger = "DiagnosticsStatusNotification"
	MessageTriggerFirmwareStatusNotification    MessageTrigger = "FirmwareStatusNotification"
	MessageTriggerHeartbeat                     MessageTrigger = "Heartbeat"
	MessageTriggerMeterValues                   MessageTrigger = "MeterValues"
	MessageTriggerStatusNotification            MessageTrigger = "StatusNotification"
)

type RegistrationStatus string

const (
	RegistrationStatusAccepted RegistrationStatus = "Accepted"
	RegistrationStatusPending  RegistrationStatus = "Pending"
	RegistrationStatusRejected RegistrationStatus = "Rejected"
)

type TriggerMessageStatus string

const (
	TriggerMessageStatusAccepted       TriggerMessageStatus = "Accepted"
	TriggerMessageStatusRejected       TriggerMessageStatus = "Rejected"
	TriggerMessageStatusNotImplemented TriggerMessageStatus = "NotImplemented"
)
