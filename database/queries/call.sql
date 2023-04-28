-- name: CreateCall :one
INSERT INTO calls (
    charge_point_id,
    req_id,
    action,
    created_at
  ) VALUES ($1, $2, $3, $4)
  RETURNING *;
  
-- name: GetCallByReqID :one
SELECT * FROM calls
  WHERE charge_point_id = $1 AND req_id = $2;
