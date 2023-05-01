package types

type DataTransferStatus string

const (
	DataTransferStatusAccepted         DataTransferStatus = "Accepted"
	DataTransferStatusRejected         DataTransferStatus = "Rejected"
	DataTransferStatusUnknownMessageId DataTransferStatus = "UnknownMessageId"
	DataTransferStatusUnknownVendorId  DataTransferStatus = "UnknownVendorId"
)

type DiagnosticsStatus string

const (
	DiagnosticsStatusIdle         DiagnosticsStatus = "Idle"
	DiagnosticsStatusUploaded     DiagnosticsStatus = "Uploaded"
	DiagnosticsStatusUploadFailed DiagnosticsStatus = "UploadFailed"
	DiagnosticsStatusUploading    DiagnosticsStatus = "Uploading"
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

type FirmwareStatus string

const (
	FirmwareStatusDownloaded         FirmwareStatus = "Downloaded"
	FirmwareStatusDownloadFailed     FirmwareStatus = "DownloadFailed"
	FirmwareStatusDownloading        FirmwareStatus = "Downloading"
	FirmwareStatusIdle               FirmwareStatus = "Idle"
	FirmwareStatusInstallationFailed FirmwareStatus = "InstallationFailed"
	FirmwareStatusInstalling         FirmwareStatus = "Installing"
	FirmwareStatusInstalled          FirmwareStatus = "Installed"
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
