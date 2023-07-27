package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func AuthMiddleware(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// If the Authorization header is empty, check for JWT in the query parameters.
		if authHeader == "" {
			authHeader = r.URL.Query().Get("jwt")
		}

		// Check if the JWT is present and starts with "Bearer ".
		if !(authHeader != "" && strings.HasPrefix(authHeader, "Bearer ")) {
			logger.Log(fmt.Errorf("authorization header missing"))
			log.Printf("authorization header missing")
			utils.ErrorResponse(w, 401, "Unauthorised request")
			return
		}

		tokenString := authHeader[7:]

		isValid, userID, err := utils.VerifyJWTToken(tokenString)

		if err != nil || !isValid {
			// If there was an error parsing the token or the token is invalid, return an error response.
			logger.Log(fmt.Errorf("error parsing token: %v", err))
			log.Printf("error parsing token: %v", err)
			utils.ErrorResponse(w, 401, "Unauthorised request")
			return
		}

		id, _ := uuid.Parse(userID)

		user, _ := apiCfg.DB.GetUserById(r.Context(), id)

		handler(w, r, user)
	}
}
