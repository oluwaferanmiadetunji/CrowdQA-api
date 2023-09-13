-- name: CreatePoll :one
INSERT INTO polls (id, created_at, updated_at, event_id, poll_type, poll_question)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetPollsByEventId :many
SELECT * FROM polls WHERE event_id = $1 ORDER BY created_at;

