-- calls
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
CREATE TYPE reservation_status AS ENUM (
    'Accepted',
    'Cancelled',
    'Completed',
	'Faulted',
	'Occupied',
	'Rejected',
	'Unavailable'
);

CREATE TABLE IF NOT EXISTS reservations (
    id              BIGSERIAL PRIMARY KEY,
    connector_id    INTEGER NOT NULL,
    charge_point_id BIGINT NOT NULL,
    req_id          TEXT NOT NULL,
    expiry_date     TIMESTAMPTZ NOT NULL,
    status          reservation_status NOT NULL,
    id_tag          TEXT NOT NULL,
    parent_id_tag   TEXT,
    created_at      TIMESTAMPTZ NOT NULL,
    updated_at      TIMESTAMPTZ NOT NULL
);

ALTER TABLE reservations 
    ADD CONSTRAINT fk_reservations_charge_point_id
    FOREIGN KEY (charge_point_id) 
    REFERENCES charge_points(id) 
    ON DELETE CASCADE;
