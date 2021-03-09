package users

import "github.com/gorilla/mux"

func SetupRouterUsers(r *mux.Router) {
	r.HandleFunc("", getUserProfile).Methods("GET", "OPTIONS")
	r.HandleFunc("", updateUserProfile).Methods("POST", "OPTIONS")
	r.HandleFunc("/{userID}", getUserProfileByID).Methods("GET", "OPTIONS")
}
