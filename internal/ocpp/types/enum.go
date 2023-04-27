package types

type Action string

const (
	ActionAUTHORIZE                     Action = "Authorize"
	ActionBOOTNOTIFICATION              Action = "BootNotification"
	ActionCANCELRESERVATION             Action = "CancelReservation"
	ActionCHANGEAVAILABILITY            Action = "ChangeAvailability"
	ActionCHANGECONFIGURATION           Action = "ChangeConfiguration"
	ActionCLEARCACHE                    Action = "ClearCache"
	ActionCLEARCHARGINGPROFILE          Action = "ClearChargingProfile"
	ActionDATATRANSFER                  Action = "DataTransfer"
	ActionDIAGNOSTICSSTATUSNOTIFICATION Action = "DiagnosticsStatusNotification"
	ActionFIRMWARESTATUSNOTIFICATION    Action = "FirmwareStatusNotification"
	ActionGETCOMPOSITESCHEDULE          Action = "GetCompositeSchedule"
	ActionGETCONFIGURATION              Action = "GetConfiguration"
	ActionGETDIAGNOSTICS                Action = "GetDiagnostics"
	ActionGETLOCALLISTVERSION           Action = "GetLocalListVersion"
	ActionHEARTBEAT                     Action = "Heartbeat"
	ActionMETERVALUES                   Action = "MeterValues"
	ActionREMOTESTARTTRANSACTION        Action = "RemoteStartTransaction"
	ActionREMOTESTOPTRANSACTION         Action = "RemoteStopTransaction"
	ActionRESERVENOW                    Action = "ReserveNow"
	ActionRESET                         Action = "Reset"
	ActionSENDLOCALLIST                 Action = "SendLocalList"
	ActionSETCHARGINGPROFILE            Action = "SetChargingProfile"
	ActionSTARTTRANSACTION              Action = "StartTransaction"
	ActionSTATUSNOTIFICATION            Action = "StatusNotification"
	ActionSTOPTRANSACTION               Action = "StopTransaction"
	ActionTRIGGERMESSAGE                Action = "TriggerMessage"
	ActionUNLOCKCONNECTOR               Action = "UnlockConnector"
	ActionUPDATEFIRMWARE                Action = "UpdateFirmware"
)

type ErrorCode string

const (
	ErrorCodeNOTIMPLEMENTED               ErrorCode = "NotImplemented"
	ErrorCodeNOTSUPPORTED                 ErrorCode = "NotSupported"
	ErrorCodeINTERNALERROR                ErrorCode = "InternalError"
	ErrorCodePROTOCOLERROR                ErrorCode = "ProtocolError"
	ErrorCodeSECURITYERROR                ErrorCode = "SecurityError"
	ErrorCodeFORMATIONVIOLATION           ErrorCode = "FormationViolation"
	ErrorCodePROPERTYCONSTRAINTVIOLATION  ErrorCode = "PropertyConstraintViolation"
	ErrorCodeOCCURENCECONSTRAINTVIOLATION ErrorCode = "OccurenceConstraintViolation"
	ErrorCodeTYPECONSTRAINTVIOLATION      ErrorCode = "TypeConstraintViolation"
	ErrorCodeGENERICERROR                 ErrorCode = "GenericError"
)

type MessageType int

const (
	MessageTypeCALL       MessageType = 2
	MessageTypeCALLRESULT MessageType = 3
	MessageTypeCALLERROR  MessageType = 4
)

type MessageTrigger string

const (
	MessageTriggerBOOTNOTIFICATION              MessageTrigger = "BootNotification"
	MessageTriggerDIAGNOSTICSSTATUSNOTIFICATION MessageTrigger = "DiagnosticsStatusNotification"
	MessageTriggerFIRMWARESTATUSNOTIFICATION    MessageTrigger = "FirmwareStatusNotification"
	MessageTriggerHEARTBEAT                     MessageTrigger = "Heartbeat"
	MessageTriggerMETERVALUES                   MessageTrigger = "MeterValues"
	MessageTriggerSTATUSNOTIFICATION            MessageTrigger = "StatusNotification"
)

type RegistrationStatus string

const (
	RegistrationStatusACCEPTED RegistrationStatus = "Accepted"
	RegistrationStatusPENDING  RegistrationStatus = "Pending"
	RegistrationStatusREJECTED RegistrationStatus = "Rejected"
)

type TriggerMessageStatus string

const (
	TriggerMessageStatusACCEPTED       TriggerMessageStatus = "Accepted"
	TriggerMessageStatusREJECTED       TriggerMessageStatus = "Rejected"
	TriggerMessageStatusNOTIMPLEMENTED TriggerMessageStatus = "NotImplemented"
)
