-- Reservations
ALTER TABLE IF EXISTS reservations 
    DROP CONSTRAINT IF EXISTS fk_reservations_charge_point_id;

DROP TABLE IF EXISTS reservations;

DROP TYPE IF EXISTS reservation_status;

-- Calls
ALTER TABLE IF EXISTS calls 
    DROP CONSTRAINT IF EXISTS fk_calls_charge_point_id;

DROP TABLE IF EXISTS calls;

DROP TYPE IF EXISTS call_action;
