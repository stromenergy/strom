-- Meter values
CREATE TYPE meter_value_format AS ENUM (
    'Raw',
	'SignedData'
);

CREATE TYPE meter_reading_context AS ENUM (
    'Interruption.Begin',
	'Interruption.End',
	'Other',
	'Sample.Clock',
	'Sample.Periodic',
	'Transaction.Begin',
	'Transaction.End',
	'Trigger'
);

CREATE TYPE meter_measurand AS ENUM (
    'Current.Export',
	'Current.Import',
	'Current.Offered',
	'Energy.Active.Export.Register',
	'Energy.Active.Import.Register',
	'Energy.Reactive.Export.Register',
	'Energy.Reactive.Import.Register',
	'Energy.Active.Export.Interval',
	'Energy.Active.Import.Interval',
	'Energy.Reactive.Export.Interval',
	'Energy.Reactive.Import.Interval',
	'Frequency',
	'Power.Active.Export',
	'Power.Active.Import',
	'Power.Factor',
	'Power.Offered',
	'Power.Reactive.Export',
	'Power.Reactive.Import',
	'RPM',
	'SoC',
	'Temperature',
	'Voltage'
);

CREATE TYPE meter_phase AS ENUM (
    'L1',
	'L2',
	'L3',
	'N',
	'L1-N',
	'L2-N',
	'L3-N',
	'L1-L2',
	'L2-L3',
	'L3-L1'
);

CREATE TYPE meter_location AS ENUM (
    'Body',
	'Cable',
	'EV',
	'Inlet',
	'Outlet'
);

CREATE TYPE meter_unit_of_measure AS ENUM (
    'Wh',
	'kWh',
	'varh',
	'kvarh',
	'W',
	'kW',
	'VA',
	'kVA',
	'var',
	'kvar',
	'A',
	'V',
	'Celsius',
	'Fahrenheit',
	'K',
	'Percent'
);

CREATE TABLE IF NOT EXISTS meter_values (
    id                BIGSERIAL PRIMARY KEY,
    connector_id      INTEGER NOT NULL,
    charge_point_id   BIGINT NOT NULL,
	transaction_id    BIGINT,
    format            meter_value_format NOT NULL,
    context           meter_reading_context NOT NULL,
    measurand         meter_measurand NOT NULL,
    phase             meter_phase,
    location          meter_location NOT NULL,
    unit              meter_unit_of_measure,
    raw_value         FLOAT,
    signed_data_value TEXT,
    timestamp         TIMESTAMPTZ NOT NULL,
    created_at        TIMESTAMPTZ NOT NULL
);

ALTER TABLE meter_values 
    ADD CONSTRAINT fk_meter_values_charge_point_id
    FOREIGN KEY (charge_point_id) 
    REFERENCES charge_points(id) 
    ON DELETE CASCADE;

ALTER TABLE meter_values 
    ADD CONSTRAINT fk_meter_values_transaction_id
    FOREIGN KEY (transaction_id) 
    REFERENCES transactions(id) 
    ON DELETE CASCADE;
