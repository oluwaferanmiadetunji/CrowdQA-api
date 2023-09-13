-- name: CreatePollOptions :one
INSERT INTO poll_options (id, created_at, updated_at, poll_id, option_text)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPollOptionsByPollId :many
SELECT * FROM poll_options WHERE poll_id = $1 ORDER BY created_at;

