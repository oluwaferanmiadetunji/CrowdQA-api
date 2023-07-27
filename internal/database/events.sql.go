// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: events.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (id, created_at, updated_at, name, start_date, end_date, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at, name, start_date, end_date, user_id
`

type CreateEventParams struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	UserID    uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.StartDate,
		arg.EndDate,
		arg.UserID,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.StartDate,
		&i.EndDate,
		&i.UserID,
	)
	return i, err
}

const getEventById = `-- name: GetEventById :one
SELECT id, created_at, updated_at, name, start_date, end_date, user_id FROM events WHERE id = $1
`

func (q *Queries) GetEventById(ctx context.Context, id uuid.UUID) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEventById, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.StartDate,
		&i.EndDate,
		&i.UserID,
	)
	return i, err
}

const getMyEvents = `-- name: GetMyEvents :many
SELECT id, created_at, updated_at, name, start_date, end_date, user_id FROM events WHERE user_id = $1
`

func (q *Queries) GetMyEvents(ctx context.Context, userID uuid.UUID) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getMyEvents, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Event{}
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.StartDate,
			&i.EndDate,
			&i.UserID,
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
