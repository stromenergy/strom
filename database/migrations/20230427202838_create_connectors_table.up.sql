-- Connectors
CREATE TYPE charge_point_error_code AS ENUM (
    'ConnectorLockFailure',
	'EVCommunicationError',
	'GroundFailure',
	'HighTemperature',
	'InternalError',
	'LocalListConflict',
	'NoError',
	'OtherError',
	'OverCurrentFailure',
	'OverVoltage',
	'PowerMeterFailure',
	'PowerSwitchFailure',
	'ReaderFailure',
	'ResetFailure',
	'UnderVoltage',
	'WeakSignal'
);

CREATE TYPE charge_point_status AS ENUM (
    'Available',
	'Preparing',
	'Charging',
	'SuspendedEVSE',
	'SuspendedEV',
	'Finishing',
	'Reserved',
	'Unavailable',
	'Faulted'
);

CREATE TABLE IF NOT EXISTS connectors (
    id                BIGSERIAL PRIMARY KEY,
    connector_id      INTEGER NOT NULL,
    charge_point_id   BIGINT NOT NULL,
    error_code        charge_point_error_code NOT NULL,
    status            charge_point_status NOT NULL,
    info              TEXT,
    vendor_id         TEXT,
    vendor_error_code TEXT,
    created_at        TIMESTAMPTZ NOT NULL,
    updated_at        TIMESTAMPTZ NOT NULL
);

ALTER TABLE connectors 
    ADD CONSTRAINT fk_connectors_charge_point_id
    FOREIGN KEY (charge_point_id) 
    REFERENCES charge_points(id) 
    ON DELETE CASCADE;
