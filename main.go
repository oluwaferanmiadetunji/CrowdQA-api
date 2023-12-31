package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/auth"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/events"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/polls"
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
	auth.AuthRoutes(route)
	events.EventRoutes(route)
	polls.PollRoutes(route)

	log.Printf("Server starting on port: %v", port)
	http.ListenAndServe(":"+port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(route))

}

func Home(w http.ResponseWriter, _ *http.Request) {
	status := utils.Response{
		Message: "Welcome",
	}

	utils.JSONResponse(w, 200, status)
}
