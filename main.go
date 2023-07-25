package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/users"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/internal/utils"
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

	route := mux.NewRouter()

	route.HandleFunc("/", Home).Methods("GET")
	users.UserRoutes(route)

	log.Printf("Server starting on port: %v", port)
	http.ListenAndServe(":"+port, route)

}

func Home(w http.ResponseWriter, r *http.Request) {
	status := utils.Response{
		Message: "Welcome",
	}

	utils.JSONResponse(w, 200, status)

}