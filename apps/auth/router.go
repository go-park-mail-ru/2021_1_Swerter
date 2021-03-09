package auth

import (
	"github.com/gorilla/mux"
)

func SetupRouterAuth(r *mux.Router) {
	r.HandleFunc("/login", login).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", logout).Methods("POST", "OPTIONS")
	r.HandleFunc("/register", register).Methods("POST", "OPTIONS")
}
