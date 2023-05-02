-- Calls
CREATE TYPE call_action AS ENUM (
    'Authorize',
	'BootNotification',
	'CancelReservation',
	'ChangeAvailability',
	'ChangeConfiguration',
	'ClearCache',
	'ClearChargingProfile',
	'DataTransfer',
	'DiagnosticsStatusNotification',
	'FirmwareStatusNotification',
	'GetCompositeSchedule',
	'GetConfiguration',
	'GetDiagnostics',
	'GetLocalListVersion',
	'Heartbeat',
	'MeterValues',
	'RemoteStartTransaction',
	'RemoteStopTransaction',
	'ReserveNow',
	'Reset',
	'SendLocalList',
	'SetChargingProfile',
	'StartTransaction',
	'StatusNotification',
	'StopTransaction',
	'TriggerMessage',
	'UnlockConnector',
	'UpdateFirmware'
);

CREATE TABLE IF NOT EXISTS calls (
    id              BIGSERIAL PRIMARY KEY,
    charge_point_id BIGINT NOT NULL,
    req_id          TEXT NOT NULL,
    action          call_action NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL
);

ALTER TABLE calls 
    ADD CONSTRAINT fk_calls_charge_point_id
    FOREIGN KEY (charge_point_id) 
    REFERENCES charge_points(id) 
    ON DELETE CASCADE;

-- Reservations
ALTER TABLE reservations
    ADD COLUMN req_id TEXT NOT NULL DEFAULT '';
