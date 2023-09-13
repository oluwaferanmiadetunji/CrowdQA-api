// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: polls.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createPoll = `-- name: CreatePoll :one
INSERT INTO polls (id, created_at, updated_at, event_id, poll_type, poll_question)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, event_id, poll_type, poll_question
`

type CreatePollParams struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	EventID      uuid.UUID `json:"event_id"`
	PollType     string    `json:"poll_type"`
	PollQuestion string    `json:"poll_question"`
}

func (q *Queries) CreatePoll(ctx context.Context, arg CreatePollParams) (Poll, error) {
	row := q.db.QueryRowContext(ctx, createPoll,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.EventID,
		arg.PollType,
		arg.PollQuestion,
	)
	var i Poll
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EventID,
		&i.PollType,
		&i.PollQuestion,
	)
	return i, err
}

const getPollsByEventId = `-- name: GetPollsByEventId :many
SELECT id, created_at, updated_at, event_id, poll_type, poll_question FROM polls WHERE event_id = $1 ORDER BY created_at
`

func (q *Queries) GetPollsByEventId(ctx context.Context, eventID uuid.UUID) ([]Poll, error) {
	rows, err := q.db.QueryContext(ctx, getPollsByEventId, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Poll{}
	for rows.Next() {
		var i Poll
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.EventID,
			&i.PollType,
			&i.PollQuestion,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
