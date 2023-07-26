package utils

import (
	"time"

	"github.com/google/uuid"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

func ConvertDatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		Name:  dbUser.Name,
		Email: dbUser.Email,
	}
}