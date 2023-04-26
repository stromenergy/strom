package ocpp

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

type MessageType int

const (
	MessageTypeCALL       MessageType = 2
	MessageTypeCALLRESULT MessageType = 3
	MessageTypeCALLERROR  MessageType = 4
)