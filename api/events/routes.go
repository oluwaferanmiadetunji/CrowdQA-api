package events

import (
	"github.com/gorilla/mux"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/middleware"
)

func EventRoutes(r *mux.Router) {
	r.HandleFunc("/events", middleware.AuthMiddleware(CreateEvents)).Methods("POST")
	r.HandleFunc("/events", middleware.AuthMiddleware(GetMyEvents)).Methods("GET")
	r.HandleFunc("/events/upcoming", middleware.AuthMiddleware(GetMyUpcomingEvents)).Methods("GET")
	r.HandleFunc("/events/{id}", middleware.AuthMiddleware(DeleteEvent)).Methods("DELETE")
	r.HandleFunc("/events/{id}", middleware.AuthMiddleware(GetMyEventById)).Methods("GET")
}
