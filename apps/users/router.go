package users

import "github.com/gorilla/mux"

func SetupRouterUsers(r *mux.Router) {
	r.HandleFunc("", userProfile).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/{userID}", getUserProfileByID).Methods("GET", "OPTIONS")
}
