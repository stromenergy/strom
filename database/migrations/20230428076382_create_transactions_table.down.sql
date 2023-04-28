-- Transactions
ALTER TABLE IF EXISTS transactions 
    DROP CONSTRAINT IF EXISTS fk_transactions_reservation_id;

ALTER TABLE IF EXISTS transactions 
    DROP CONSTRAINT IF EXISTS fk_transactions_charge_point_id;

DROP TABLE IF EXISTS transactions;

DROP TYPE IF EXISTS transaction_stop_reason;

DROP TYPE IF EXISTS transaction_status;
