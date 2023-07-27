package events

import (
	"github.com/gorilla/mux"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/middleware"
)

func EventRoutes(r *mux.Router) {
	r.HandleFunc("/events", middleware.AuthMiddleware(CreateEvents)).Methods("POST")

}
