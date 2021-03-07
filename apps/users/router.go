package users

import "github.com/gorilla/mux"

func SetupRouterUsers(r *mux.Router) {
	r.HandleFunc("", getUserProfile).Methods("GET")
	r.HandleFunc("", updateUserProfile).Methods("POST")
	r.HandleFunc("/{userID}", getUserByID).Methods("GET")
}
