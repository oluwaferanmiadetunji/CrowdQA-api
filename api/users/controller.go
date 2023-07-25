package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/db"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

var (
	apiCfg    = db.Database()
	logger, _ = thoth.Init("log")
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		logger.Log(fmt.Errorf("error parsing data %v", err))
		log.Printf("error parsing data: %v", err)
		utils.ErrorResponse(w, 400, "Error creating account, please try again")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  params.Password,
	})

	if err != nil {
		logger.Log(fmt.Errorf("error creating account %v", err))
		log.Printf("error creating account: %v", err)
		utils.ErrorResponse(w, 400, ("Error creating account, please try again"))
		return
	}

	utils.JSONResponse(w, 201, ConvertDatabaseUserToUser(user))
}
