-- name: CreateIDTag :one
INSERT INTO id_tags (
    parent_id_tag_id,
    token, 
    status,
    created_at,
    updated_at
  ) VALUES ($1, $2, $3, $4, $5)
  RETURNING *;

-- name: GetIDTag :one
SELECT * FROM id_tags
  WHERE id = $1;

-- name: GetIDTagByToken :one
SELECT * FROM id_tags
  WHERE token = $1;

-- name: ListIDTags :many
SELECT * FROM id_tags
  ORDER BY id;

-- name: UpdateIDTag :one
UPDATE id_tags SET (
    parent_id_tag_id,
    status,
    updated_at
  ) = ($2, $3, $4)
  WHERE id = $1
  RETURNING *;
