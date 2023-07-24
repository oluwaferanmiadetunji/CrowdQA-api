package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api"
)

func main() {
	logger, _ := thoth.Init("log")

	if err := godotenv.Load(); err != nil {
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	log.Printf("Server starting on port: %v", port)
	err := http.ListenAndServe(":"+port, api.Init())

	if err != nil {
		log.Fatal(err)
	}

}
