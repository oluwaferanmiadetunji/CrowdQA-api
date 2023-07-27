-- name: CreateEvent :one
INSERT INTO events (id, created_at, updated_at, name, start_date, end_date, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetEventById :one
SELECT * FROM events WHERE id = $1;

-- name: GetMyEvents :many
SELECT * FROM events WHERE user_id = $1;