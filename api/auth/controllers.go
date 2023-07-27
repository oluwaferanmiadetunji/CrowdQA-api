package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/db"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

var (
	apiCfg    = db.Database()
	logger, _ = thoth.Init("log")
)

func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		logger.Log(fmt.Errorf("error parsing data %v", err))
		log.Printf("error parsing data: %v", err)
		utils.ErrorResponse(w, 400, "Error logging in, please try again")
		return
	}

	existingUser, _ := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)

	if existingUser.Email == "" {
		logger.Log(fmt.Errorf("account does not exist "))
		log.Printf("account does not exist")
		utils.ErrorResponse(w, 404, "Invalid credentials!")
		return
	}

	isPasswordCorrect := utils.ComparePasswordHash(existingUser.Password, []byte(params.Password))

	if !isPasswordCorrect {
		logger.Log(fmt.Errorf("invalid credentials "))
		log.Printf("invalid credentials")
		utils.ErrorResponse(w, 417, "Invalid credentials!")
		return
	}

	token, err := utils.GenerateJWTToken(existingUser)

	if err != nil {
		logger.Log(fmt.Errorf("error generating token %v", err))
		log.Printf("error generating token: %v", err)
		utils.ErrorResponse(w, 400, "Error logging in, please try again")
		return
	}

	utils.JSONResponse(w, 200, utils.ReturnTokenResponse(token, utils.ConvertDatabaseUserToUser(existingUser)))
}
