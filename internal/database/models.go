// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	UserID    uuid.UUID `json:"user_id"`
	EventCode int32     `json:"event_code"`
}

type Poll struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	EventID      uuid.UUID `json:"event_id"`
	PollType     string    `json:"poll_type"`
	PollQuestion string    `json:"poll_question"`
}

type PollOption struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	PollID     uuid.UUID `json:"poll_id"`
	OptionText string    `json:"option_text"`
}

type PollResponse struct {
	ID             uuid.UUID      `json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	PollID         uuid.UUID      `json:"poll_id"`
	UserID         sql.NullString `json:"user_id"`
	SelectedOption sql.NullInt32  `json:"selected_option"`
	UserAnswer     sql.NullString `json:"user_answer"`
	UserRanking    []int32        `json:"user_ranking"`
	UserRating     sql.NullInt32  `json:"user_rating"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}
