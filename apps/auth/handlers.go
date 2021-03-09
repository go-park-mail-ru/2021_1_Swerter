package auth

import (
	"encoding/json"
	"fmt"
	i "my-motivation/internal"
	"my-motivation/utils"
	"net/http"
	"time"
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

	decoder := json.NewDecoder(r.Body)
	user := i.User{}
	decoder.Decode(&user)
	fmt.Printf("%+v\n", user)

	if isUserExist(user) {
		expiration := time.Now().Add(10 * time.Hour)
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    user.Login,
			Expires:  expiration,
		}
		i.SessionsCounter++
		i.Sessions[user.Login] = user
		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(loginSuccess))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	session, err := r.Cookie("session_id")

	if err == http.ErrNoCookie {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(logoutSuccess))
}

func register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	decoder := json.NewDecoder(r.Body)
	newUser := i.User{}
	decoder.Decode(&newUser)
	i.IDCounter++
	newUser.ID = i.IDCounter

	i.Users[newUser.Login] = newUser

	fmt.Printf("%+v\n", newUser)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(registerSuccess))
}

func isUserExist(user i.User) bool {
	if u, ok := i.Users[user.Login]; ok {
		if u.Password == user.Password {
			return true
		}
	}
	return false
}
