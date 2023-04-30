-- ID tags
ALTER TABLE IF EXISTS id_tags 
    DROP CONSTRAINT IF EXISTS fk_id_tags_parent_id_tag_id;

DROP TABLE IF EXISTS id_tags;

DROP TYPE IF EXISTS authorization_status;
