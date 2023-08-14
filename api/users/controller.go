package users

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

type UserParameters struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := UserParameters{}

	err := decoder.Decode(&params)

	if err != nil {
		logger.Log(fmt.Errorf("error parsing data %v", err))
		log.Printf("error parsing data: %v", err)
		utils.ErrorResponse(w, 400, "Error creating account, please try again")
		return
	}

	existingUser, _ := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)

	if existingUser.Email != "" {
		utils.ErrorResponse(w, 409, "Account already created!")
		return
	}

	user, err := SaveUserToDb(params)

	if err != nil {
		logger.Log(fmt.Errorf("error creating account %v", err))
		log.Printf("error creating account: %v", err)
		utils.ErrorResponse(w, 400, ("Error creating account, please try again"))
		return
	}

	token, err := utils.GenerateJWTToken(existingUser)

	if err != nil {
		logger.Log(fmt.Errorf("error generating token %v", err))
		log.Printf("error generating token: %v", err)
		utils.ErrorResponse(w, 400, "Error logging in, please try again")
		return
	}

	utils.JSONResponse(w, 201, utils.ReturnTokenResponse(token, utils.ConvertDatabaseUserToUser(*user)))
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {}
