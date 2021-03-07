package auth

import (
	"github.com/gorilla/mux"
)

func SetupRouterAuth(r *mux.Router) {
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")
}
