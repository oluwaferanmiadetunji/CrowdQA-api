package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/ichtrojan/thoth"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) []byte {
	logger, _ := thoth.Init("log")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		logger.Log(fmt.Errorf("error hashing password %v", err))
		log.Printf("error hashing password: %v", err)
	}
	return hashedPassword

}

func ComparePasswordHash(hashedPassword string, password []byte) bool {
	logger, _ := thoth.Init("log")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)

	if err != nil {
		logger.Log(fmt.Errorf("passwords are not the same %v", err))
		log.Printf("passwords are not the same: %v", err)
		return false
	}

	return true
}

func ConvertStringToTime(date string) time.Time {
	parsedTime, _ := time.Parse(time.RFC3339, date)
	return parsedTime
}
