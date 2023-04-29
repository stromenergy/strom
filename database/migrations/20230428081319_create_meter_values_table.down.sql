-- Meter values
ALTER TABLE IF EXISTS meter_values 
    DROP CONSTRAINT IF EXISTS fk_meter_values_transaction_id;

ALTER TABLE IF EXISTS meter_values 
    DROP CONSTRAINT IF EXISTS fk_meter_values_charge_point_id;

DROP TABLE IF EXISTS meter_values;

DROP TYPE IF EXISTS meter_unit_of_measure;

DROP TYPE IF EXISTS meter_location;

DROP TYPE IF EXISTS meter_phase;

DROP TYPE IF EXISTS meter_measurand;

DROP TYPE IF EXISTS meter_reading_context;

DROP TYPE IF EXISTS meter_value_format;
