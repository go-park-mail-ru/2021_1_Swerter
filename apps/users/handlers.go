package users

import (
	"net/http"
	"my-motivation/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func getUserProfile(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", 401)
		return
	}

	if isSessionExist(session.Value) {
		user := i.Sessions[session.Value]
		userJson, _ := json.Marshal(&user)
		w.Write(userJson)
	}

	w.Write([]byte("getUserProfile"))
}

func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("updateUserProfile"))
}

func getUserProfileByID(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("getUserByID" + " " + mux.Vars(r)["userID"]))
}

func isSessionExist(session string) bool {
	if _, ok := i.Sessions[session]; ok {
		return true
	}
	return false
}
