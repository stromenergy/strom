package types

type AvailabilityStatus string

const (
	AvailabilityStatusAccepted  AvailabilityStatus = "Accepted"
	AvailabilityStatusRejected  AvailabilityStatus = "Rejected"
	AvailabilityStatusScheduled AvailabilityStatus = "Scheduled"
)

type AvailabilityType string

const (
	AvailabilityTypeInoperative AvailabilityType = "Inoperative"
	AvailabilityTypeOperative   AvailabilityType = "Operative"
)

type CallAction string

const (
	CallActionAuthorize                     CallAction = "Authorize"
	CallActionBootNotification              CallAction = "BootNotification"
	CallActionCancelReservation             CallAction = "CancelReservation"
	CallActionChangeAvailability            CallAction = "ChangeAvailability"
	CallActionChangeConfiguration           CallAction = "ChangeConfiguration"
	CallActionClearCache                    CallAction = "ClearCache"
	CallActionClearChargingProfile          CallAction = "ClearChargingProfile"
	CallActionDataTransfer                  CallAction = "DataTransfer"
	CallActionDiagnosticsStatusNotification CallAction = "DiagnosticsStatusNotification"
	CallActionFirmwareStatusNotification    CallAction = "FirmwareStatusNotification"
	CallActionGetCompositeSchedule          CallAction = "GetCompositeSchedule"
	CallActionGetConfiguration              CallAction = "GetConfiguration"
	CallActionGetDiagnostics                CallAction = "GetDiagnostics"
	CallActionGetLocalListVersion           CallAction = "GetLocalListVersion"
	CallActionHeartbeat                     CallAction = "Heartbeat"
	CallActionMeterValues                   CallAction = "MeterValues"
	CallActionRemoteStartTransaction        CallAction = "RemoteStartTransaction"
	CallActionRemoteStopTransaction         CallAction = "RemoteStopTransaction"
	CallActionReserveNow                    CallAction = "ReserveNow"
	CallActionReset                         CallAction = "Reset"
	CallActionSendLocalList                 CallAction = "SendLocalList"
	CallActionSetChargingProfile            CallAction = "SetChargingProfile"
	CallActionStartTransaction              CallAction = "StartTransaction"
	CallActionStatusNotification            CallAction = "StatusNotification"
	CallActionStopTransaction               CallAction = "StopTransaction"
	CallActionTriggerMessage                CallAction = "TriggerMessage"
	CallActionUnlockConnector               CallAction = "UnlockConnector"
	CallActionUpdateFirmware                CallAction = "UpdateFirmware"
)

type CancelReservationStatus string

const (
	CancelReservationStatusAccepted CancelReservationStatus = "Accepted"
	CancelReservationStatusRejected CancelReservationStatus = "Rejected"
)

type ConfigurationStatus string

const (
	ConfigurationStatusAccepted       ConfigurationStatus = "Accepted"
	ConfigurationStatusRejected       ConfigurationStatus = "Rejected"
	ConfigurationStatusRebootRequired ConfigurationStatus = "RebootRequired"
	ConfigurationStatusNotSupported   ConfigurationStatus = "NotSupported"
)

type ClearCacheStatus string

const (
	ClearCacheStatusAccepted ClearCacheStatus = "Accepted"
	ClearCacheStatusRejected ClearCacheStatus = "Rejected"
)

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

type RemoteStartStopStatus string

const (
	RemoteStartStopStatusAccepted RemoteStartStopStatus = "Accepted"
	RemoteStartStopStatusRejected RemoteStartStopStatus = "Rejected"
)

type ResetStatus string

const (
	ResetStatusAccepted ResetStatus = "Accepted"
	ResetStatusRejected ResetStatus = "Rejected"
)

type ResetType string

const (
	ResetTypeHard ResetType = "Hard"
	ResetTypeSoft ResetType = "Soft"
)

type TriggerMessageStatus string

const (
	TriggerMessageStatusAccepted       TriggerMessageStatus = "Accepted"
	TriggerMessageStatusRejected       TriggerMessageStatus = "Rejected"
	TriggerMessageStatusNotImplemented TriggerMessageStatus = "NotImplemented"
)

type UnlockStatus string

const (
	UnlockStatusUnlocked     UnlockStatus = "Unlocked"
	UnlockStatusUnlockFailed UnlockStatus = "UnlockFailed"
	UnlockStatusNotSupported UnlockStatus = "NotSupported"
)
