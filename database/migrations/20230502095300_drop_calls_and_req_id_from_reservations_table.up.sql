-- Reservations
ALTER TABLE reservations
    DROP IF EXISTS req_id;

-- Calls
ALTER TABLE IF EXISTS calls 
    DROP CONSTRAINT IF EXISTS fk_calls_charge_point_id;

DROP TABLE IF EXISTS calls;

DROP TYPE IF EXISTS call_action;
