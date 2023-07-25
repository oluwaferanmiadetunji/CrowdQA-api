package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/ichtrojan/thoth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/config"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/database"
)

func Database() config.ApiConfig {
	logger, _ := thoth.Init("log")

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
