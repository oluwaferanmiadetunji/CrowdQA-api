package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

func TokenExpirationTime() time.Time {
	return time.Now().Add(120 * time.Minute)
}

func GenerateJWTToken(user database.User) (string, error) {
	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	secret, exist := os.LookupEnv("API_SECRET")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	var jwtKey = []byte(secret)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = TokenExpirationTime()
	claims["authorized"] = true
	claims["id"] = user.ID

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {

		}
	})
}
