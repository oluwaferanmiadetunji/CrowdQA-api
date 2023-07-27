package events

import (
	"github.com/gorilla/mux"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/middleware"
)

func EventRoutes(r *mux.Router) {

	protectedRoute := r.PathPrefix("/events").Subrouter()

	protectedRoute.HandleFunc("", CreateEvents).Methods("POST")
	protectedRoute.Use(middleware.JWTMiddleware)

}
