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
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if isSessionExist(session.Value) {
		userID := i.Sessions[session.Value]
		user := i.Users[i.IDToLogin[userID]]

		userJson, _ := json.Marshal(&user)
		w.Write(userJson)
	}
}

func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)

	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	decoder := json.NewDecoder(r.Body)
	newUser := i.User{}
	decoder.Decode(&newUser)

	if isSessionExist(session.Value) {
		userID := i.Sessions[session.Value]
		oldUser := i.Users[i.IDToLogin[userID]]
		updateUser(&newUser, &oldUser)

		log.Printf("User update success: %+v\n", newUser)
		return
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

func updateUser(newUser *i.User, oldUser *i.User) {
	newUser.ID = oldUser.ID

	if newUser.Login == "" {
		newUser.Login = oldUser.Login
	} else {
		i.IDToLogin[newUser.ID] = newUser.Login
	}

	if newUser.Password == "" {
		newUser.Password = oldUser.Password
	} else {
		newUser.Password = utils.HashPassword(newUser.Password)
	}

	if newUser.FirstName == "" {
		newUser.FirstName = oldUser.FirstName
	}

	if newUser.LastName == "" {
		newUser.LastName = oldUser.LastName
	}

	delete(i.Users, oldUser.Login)
	i.Users[newUser.Login] = *newUser
}

