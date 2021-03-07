package users

import (
	"net/http"

	"../../utils"
	"github.com/gorilla/mux"
)

func getUserProfile(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("getUserProfile"))
}

func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("updateUserProfile"))
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("getUserByID" + " " + mux.Vars(r)["userID"]))
}
