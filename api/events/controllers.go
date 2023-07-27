package events

import (
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

func CreateEvents(w http.ResponseWriter, r *http.Request) {
	userID := utils.HandleNoUserIDResponse(w, r)

	existingUser, _ := apiCfg.DB.GetUserById(r.Context(), userID)

	if existingUser.Email == "" {
		logger.Log(fmt.Errorf("account does not exist "))
		log.Printf("account does not exist")
		utils.ErrorResponse(w, 404, "Unauthorised request!")
		return
	}

	utils.JSONResponse(w, 200, existingUser)

}
