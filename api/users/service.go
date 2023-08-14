package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

func SaveUserToDb(params UserParameters) (*database.User, error) {
	ctx := context.Background()
	user, err := apiCfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(utils.HashPassword(params.Password)),
	})

	if err != nil {
		logger.Log(fmt.Errorf("error creating account %v", err))
		log.Printf("error creating account: %v", err)

		return nil, err
	}

	return &user, nil
}
