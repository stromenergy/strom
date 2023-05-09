-- Charge points
ALTER TABLE charge_points 
    DROP password;
    
ALTER TABLE charge_points 
    DROP status;

DROP TYPE IF EXISTS charge_point_status;

ALTER TYPE connector_status RENAME TO charge_point_status;
