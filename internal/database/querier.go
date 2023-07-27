// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetEventById(ctx context.Context, id uuid.UUID) (Event, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (User, error)
}

var _ Querier = (*Queries)(nil)
