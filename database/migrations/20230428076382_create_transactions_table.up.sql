-- Transactions
CREATE TYPE transaction_status AS ENUM (
    'Started',
	'Stopped'
);

CREATE TYPE transaction_stop_reason AS ENUM (
    'EmergencyStop',
	'EVDisconnected',
	'HardReset',
	'Other',
	'Outlet',
	'PowerLoss',
	'Reboot',
	'Remote',
	'SoftReset',
	'UnlockCommand',
	'DeAuthorized'
);

CREATE TABLE IF NOT EXISTS transactions (
    id              BIGSERIAL PRIMARY KEY,
    connector_id    INTEGER NOT NULL,
    charge_point_id BIGINT NOT NULL,
    reservation_id  BIGINT,
    status          transaction_status NOT NULL DEFAULT 'Started',
    id_tag          TEXT NOT NULL,
    reason          transaction_stop_reason NULL,
    meter_start     INTEGER NOT NULL,
    meter_stop      INTEGER,
    start_timestamp TIMESTAMPTZ NOT NULL,
    stop_timestamp  TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL,
    updated_at      TIMESTAMPTZ NOT NULL
);

ALTER TABLE transactions 
    ADD CONSTRAINT fk_transactions_charge_point_id
    FOREIGN KEY (charge_point_id) 
    REFERENCES charge_points(id) 
    ON DELETE CASCADE;

ALTER TABLE transactions 
    ADD CONSTRAINT fk_transactions_reservation_id
    FOREIGN KEY (reservation_id) 
    REFERENCES reservations(id) 
    ON DELETE CASCADE;
