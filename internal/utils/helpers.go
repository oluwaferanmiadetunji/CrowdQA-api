package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
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

func GenerateEventCode() (int32, error) {
	const charset = "0123456789"
	randomNumber := make([]byte, 6)

	_, err := rand.Read(randomNumber)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(randomNumber); i++ {
		randomNumber[i] = charset[randomNumber[i]%byte(len(charset))]
	}

	// Convert the random number string to a big integer
	randomInt, success := new(big.Int).SetString(string(randomNumber), 10)
	if !success {
		return 0, fmt.Errorf("failed to convert random number to integer")
	}

	// Return the big integer as an integer
	return int32(randomInt.Int64()), nil
}
