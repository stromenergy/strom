-- Configurations
ALTER TABLE IF EXISTS configurations 
    DROP CONSTRAINT IF EXISTS fk_configurations_charge_point_id;

DROP TABLE IF EXISTS configurations;
