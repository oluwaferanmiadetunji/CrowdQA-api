package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"errors"
	"log"
	"os"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/config"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

func Database() config.ApiConfig {
	logger, _ := thoth.Init("log")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		logger.Log(errors.New("no .env file found"))
	}

	dbURL, exist := os.LookupEnv("DB_URL")

	if !exist {
		logger.Log(errors.New("DB_URL  not set in .env"))
		log.Fatal("DB_URL  not set in .env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		logger.Log(errors.New("error connecting to Postgres DB"))
		log.Fatal(err)
	}

	apiCfg := config.ApiConfig{
		DB: database.New(conn),
	}

	return apiCfg
}
