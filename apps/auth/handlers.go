package auth

import (
	"encoding/json"
	"fmt"
	i "my-motivation/internal"
	"my-motivation/utils"
	"net/http"
	"time"
	"log"
)

var registerSuccess string = `"status":"true"`
var registerFail string = `"status":"false"`
var loginSuccess string = `{"status":"true"}`
var loginFail string = `{"status":"false"}`
var logoutSuccess string = `{"status":"true"}`
var inProfile string = `{"Auth":"success"}`

func login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	user := i.User{}
	decoder.Decode(&user)

	if isUserExist(user) {
		log.Printf("User login success: %+v\n", user)
		expiration := time.Now().Add(10 * time.Hour)
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    user.Login,
			Expires:  expiration,
		}
		i.SessionsCounter++
		i.Sessions[user.Login] = user
		http.SetCookie(w, &cookie)
	} else {
		log.Printf("User login failed: %+v\n", user)
		w.WriteHeader(http.StatusForbidden)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)

	if r.Method == "OPTIONS" {
		return
	}

	session, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		log.Println("No cookie was provided for logout")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	log.Println("Logout")
}

func register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	newUser := i.User{}
	decoder.Decode(&newUser)

	if _, ok := i.Users[newUser.Login]; ok {
		log.Println("User already exists on register")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	i.IDCounter++
	newUser.ID = "id"+fmt.Sprint(i.IDCounter)

	i.Users[newUser.Login] = newUser

	responseBody := []byte("{\"userID\":"+newUser.ID+"}")
	w.Write(responseBody)

	fmt.Printf("New user. Private user data: %+v\n", newUser)
}

func isUserExist(user i.User) bool {
	if u, ok := i.Users[user.Login]; ok {
		if u.Password == user.Password {
			return true
		}
	}
	return false
}
