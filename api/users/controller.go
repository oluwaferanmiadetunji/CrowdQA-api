package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/db"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

var (
	apiCfg = db.Database()
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		utils.ErrorResponse(w, 400, fmt.Sprintf("Error parsing data %v", err))
	}

	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
	})

	status := utils.Response{
		Message: "Welcome",
	}

	utils.JSONResponse(w, 200, status)
}
