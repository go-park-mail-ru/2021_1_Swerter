package auth

import (
	"net/http"

	"../../utils"
)

func login(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("LOGIN"))
}

func logout(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("LOGOUT"))
}

func register(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("REGISTER"))
}
