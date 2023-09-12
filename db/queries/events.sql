-- name: CreateEvent :one
INSERT INTO events (id, created_at, updated_at, name, start_date, end_date, user_id, event_code)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetEventById :one
SELECT * FROM events WHERE id = $1 AND user_id = $2;

-- name: GetMyEvents :many
SELECT * FROM events WHERE user_id = $1 AND end_date < CURRENT_DATE ORDER BY start_date LIMIT 10 OFFSET $2;

-- name: GetMyEventsCount :one
SELECT COUNT(*) FROM events WHERE user_id = $1 AND end_date < CURRENT_DATE;

-- name: GetEventByEventCode :one
SELECT * FROM events WHERE event_code = $1;

-- name: GetUpcomingEvents :many
SELECT * FROM events WHERE user_id = $1 AND end_date >= CURRENT_DATE ORDER BY start_date LIMIT 10 OFFSET $2;

-- name: GetUpComingEventsCount :one
SELECT COUNT(*) FROM events WHERE user_id = $1 AND end_date >= CURRENT_DATE;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1 AND user_id = $2;