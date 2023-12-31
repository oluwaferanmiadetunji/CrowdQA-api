package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

func TokenExpirationTime() int64 {
	time := time.Now().Add(1440 * time.Minute).Unix()

	return time
}

func GenerateJWTToken(user database.User) (string, error) {
	logger, _ := thoth.Init("log")

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
	claims["iat"] = time.Now().Unix()
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

func VerifyJWTToken(tokenString string) (bool, string, error) {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	secret, exist := os.LookupEnv("API_SECRET")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Access the claims or any user-related information from the JWT.
		// For example:
		userID := claims["id"].(string)
		expirationTime := int64(claims["exp"].(float64))

		// Check if the token is expired.
		if time.Now().Unix() > expirationTime {
			return false, "", fmt.Errorf("token has expired")
		}

		// The token is valid.
		return true, userID, nil
	}

	return false, "", fmt.Errorf("invalid token")
}
