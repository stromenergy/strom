-- ID tags
CREATE TYPE authorization_status AS ENUM (
    'Accepted',
	'Blocked',
	'Expired',
	'Invalid',
	'ConcurrentTx'
);

CREATE TABLE IF NOT EXISTS id_tags (
    id               BIGSERIAL PRIMARY KEY,
    parent_id_tag_id BIGINT,
    token            TEXT NOT NULL,
    status           authorization_status NOT NULL DEFAULT 'Accepted',
    created_at       TIMESTAMPTZ NOT NULL,
    updated_at       TIMESTAMPTZ NOT NULL
);

ALTER TABLE id_tags 
    ADD CONSTRAINT fk_id_tags_parent_id_tag_id
    FOREIGN KEY (parent_id_tag_id) 
    REFERENCES id_tags(id) 
    ON DELETE SET NULL;
