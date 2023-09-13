package polls

import (
	"github.com/gorilla/mux"
	"github.com/oluwaferanmiadetunji/CrowdQA-api/api/middleware"
)

func PollRoutes(r *mux.Router) {
	r.HandleFunc("/polls/{id}", middleware.AuthMiddleware(CreatePoll)).Methods("POST")
	r.HandleFunc("/polls/{id}", middleware.AuthMiddleware(FetchPolls)).Methods("GET")
}
