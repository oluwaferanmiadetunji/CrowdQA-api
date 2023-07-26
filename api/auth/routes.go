package auth

import (
	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/login", GenerateTokenHandler).Methods("POST")

}
