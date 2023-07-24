package api

import (
	"encoding/json"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseJSON, err := json.Marshal(status)

	if err != nil {
		http.Error(w, "Failed to generate response JSON", http.StatusInternalServerError)
		return
	}

	w.Write(responseJSON)

}
