package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

func ConvertDatabaseUserToUser(dbUser database.User) User {

	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
	}

}
