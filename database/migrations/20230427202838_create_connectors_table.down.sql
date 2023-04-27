-- Connectors
ALTER TABLE IF EXISTS connectors 
    DROP CONSTRAINT IF EXISTS fk_connectors_charge_point_id;

DROP TABLE IF EXISTS connectors;

DROP TYPE IF EXISTS charge_point_status;

DROP TYPE IF EXISTS charge_point_error_code;
