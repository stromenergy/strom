-- Charge points
CREATE TABLE IF NOT EXISTS charge_points (
    id                  BIGSERIAL PRIMARY KEY,
    model               TEXT NOT NULL,
    vendor              TEXT NOT NULL,
    serial_number       TEXT,
    firmware_verion     TEXT,
    modem_iccid         TEXT,
    modem_imsi          TEXT,
    meter_serial_number TEXT,
    meter_type          TEXT,
    created_at          TIMESTAMPTZ NOT NULL,
    updated_at          TIMESTAMPTZ NOT NULL
);
