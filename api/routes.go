package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", Home).Methods("GET")

	return route
}

func Home(w http.ResponseWriter, r *http.Request) {
	status := Response{
		Message: "Welcome",
	}

	JSONResponse(w, 200, status)

}
