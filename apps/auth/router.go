package auth

import (
	m "my-motivation/apps/middleware"
	"github.com/gorilla/mux"
)

func SetupRouterAuth(r *mux.Router) {
	r.HandleFunc("/login", m.CORS(login)).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", m.CORS(logout)).Methods("POST", "OPTIONS")
	r.HandleFunc("/register", m.CORS(register)).Methods("POST", "OPTIONS")
}
