-- name: CreatePollResponse :one
INSERT INTO poll_responses (id, created_at, updated_at, poll_id, user_id, selected_option, user_answer, user_ranking, user_rating)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetPollResponsesByPollId :many
SELECT * FROM poll_responses WHERE poll_id = $1 ORDER BY created_at;