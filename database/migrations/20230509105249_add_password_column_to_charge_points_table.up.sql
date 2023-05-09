-- Charge points
ALTER TYPE charge_point_status RENAME TO connector_status;

CREATE TYPE charge_point_status AS ENUM (
    'Online',
	'Pending',
	'Offline'
);

ALTER TABLE charge_points 
    ADD COLUMN status charge_point_status NOT NULL DEFAULT 'Offline';

ALTER TABLE charge_points 
    ADD COLUMN password BYTEA;
