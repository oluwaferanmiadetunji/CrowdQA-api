package users

import (
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")

}
