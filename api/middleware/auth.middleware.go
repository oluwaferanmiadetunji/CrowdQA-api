package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
)

const jwtContextKey = "user"

func JWTMiddleware(next http.Handler) http.Handler {
	logger, _ := thoth.Init("log")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request.
		authHeader := r.Header.Get("Authorization")

		// If the Authorization header is empty, check for JWT in the query parameters.
		if authHeader == "" {
			authHeader = r.URL.Query().Get("jwt")
		}

		// Check if the JWT is present and starts with "Bearer ".
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			// Extract the token from the "Bearer " prefix.
			tokenString := authHeader[7:]

			isValid, userID, err := utils.VerifyJWTToken(tokenString)

			if err != nil || !isValid {
				// If there was an error parsing the token or the token is invalid, return an error response.
				logger.Log(fmt.Errorf("error parsing token: %v", err))
				log.Printf("error parsing token: %v", err)
				utils.ErrorResponse(w, 401, "Unauthorised request")
				return
			}

			ctx := context.WithValue(r.Context(), jwtContextKey, userID)
			r = r.WithContext(ctx)

		} else {
			// If the Authorization header or "jwt" query parameter is missing, return an error.
			logger.Log(fmt.Errorf("authorization header missing"))
			log.Printf("authorization header missing")
			utils.ErrorResponse(w, 401, "Unauthorised request")
			return
		}

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
