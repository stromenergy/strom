-- Configurations
CREATE TABLE IF NOT EXISTS configurations (
    id              BIGSERIAL PRIMARY KEY,
    charge_point_id BIGINT NOT NULL,
    key             TEXT NOT NULL,
    readonly        BOOLEAN NOT NULL,
    value           TEXT,
    created_at      TIMESTAMPTZ NOT NULL,
    updated_at      TIMESTAMPTZ NOT NULL
);

ALTER TABLE configurations 
    ADD CONSTRAINT fk_configurations_charge_point_id
    FOREIGN KEY (charge_point_id) 
    REFERENCES charge_points(id) 
    ON DELETE CASCADE;
