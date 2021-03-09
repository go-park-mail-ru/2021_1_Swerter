package auth

import (
	"encoding/json"
	"errors"
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

	u, err := getUser(user);
	if err!=nil {
		log.Printf("User login failed: %+v\n", user)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Printf("User login success: %+v\n", u)
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    u.Login,
		Expires:  expiration,
	}
	i.SessionsCounter++
	i.Sessions[u.Login] = u
	http.SetCookie(w, &cookie)


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

func getUser(user i.User) (i.User, error) {
	if u, ok := i.Users[user.Login]; ok {
		if u.Password == user.Password {
			return u, nil
		}
	}
	return user, errors.New("No user")
}
