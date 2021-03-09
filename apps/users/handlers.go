package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	i "my-motivation/internal"
	"my-motivation/utils"
	"net/http"
)

func userProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		getUserProfile(w, r)
	}
	if r.Method == http.MethodPost {
		updateUserProfile(w, r)
	}
}


func getUserProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if isSessionExist(session.Value) {
		user := i.Sessions[session.Value]
		userJson, _ := json.Marshal(&user)
		w.Write(userJson)
	}
}

func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	user := i.User{}
	decoder.Decode(&user)

	if isSessionExist(session.Value) {
		id := i.Sessions[session.Value].ID
		user.ID = id
		i.Sessions[session.Value] = user
		w.WriteHeader(http.StatusOK)
		log.Printf("User update success: %+v\n", user)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
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
