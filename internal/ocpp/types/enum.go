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

type ChargePointErrorCode string

const (
	ChargePointErrorCodeCONNECTORLOCKFAILURE ChargePointErrorCode = "ConnectorLockFailure"
	ChargePointErrorCodeEVCOMMUNICATIONERROR ChargePointErrorCode = "EVCommunicationError"
	ChargePointErrorCodeGROUNDFAILURE        ChargePointErrorCode = "GroundFailure"
	ChargePointErrorCodeHIGHTEMPERATURE      ChargePointErrorCode = "HighTemperature"
	ChargePointErrorCodeINTERNALERROR        ChargePointErrorCode = "InternalError"
	ChargePointErrorCodeLOCALLISTCONFLICT    ChargePointErrorCode = "LocalListConflict"
	ChargePointErrorCodeNOERROR              ChargePointErrorCode = "NoError"
	ChargePointErrorCodeOTHERERROR           ChargePointErrorCode = "OtherError"
	ChargePointErrorCodeOVERCURRENTFAILURE   ChargePointErrorCode = "OverCurrentFailure"
	ChargePointErrorCodeOVERVOLTAGE          ChargePointErrorCode = "OverVoltage"
	ChargePointErrorCodePOWERMETERFAILURE    ChargePointErrorCode = "PowerMeterFailure"
	ChargePointErrorCodePOWERSWITCHFAILURE   ChargePointErrorCode = "PowerSwitchFailure"
	ChargePointErrorCodeREADERFAILURE        ChargePointErrorCode = "ReaderFailure"
	ChargePointErrorCodeRESETFAILURE         ChargePointErrorCode = "ResetFailure"
	ChargePointErrorCodeUNDERVOLTAGE         ChargePointErrorCode = "UnderVoltage"
	ChargePointErrorCodeWEAKSIGNAL           ChargePointErrorCode = "WeakSignal"
)

type ChargePointStatus string

const (
	ChargePointStatusAVAILABLE     ChargePointStatus = "Available"
	ChargePointStatusPREPARING     ChargePointStatus = "Preparing"
	ChargePointStatusCHARGING      ChargePointStatus = "Charging"
	ChargePointStatusSUSPENDEDEVSE ChargePointStatus = "SuspendedEVSE"
	ChargePointStatusSUSPENDEDEV   ChargePointStatus = "SuspendedEV"
	ChargePointStatusFINISHING     ChargePointStatus = "Finishing"
	ChargePointStatusRESERVED      ChargePointStatus = "Reserved"
	ChargePointStatusUNAVAILABLE   ChargePointStatus = "Unavailable"
	ChargePointStatusFAULTED       ChargePointStatus = "Faulted"
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
