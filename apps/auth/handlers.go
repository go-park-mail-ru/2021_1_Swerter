package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	i "my-motivation/internal"
	"my-motivation/utils"
	"net/http"
	"time"
)

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

	u, err := getUser(user)
	if err != nil {
		log.Printf("User login failed: %+v\n", user)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Printf("User login success: %+v\n", u)
	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:    "session_id",
		Value:   utils.GenSession(u.ID),
		Expires: expiration,
	}
	i.SessionsCounter++
	i.Sessions[cookie.Value] = u.ID
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
	log.Println("Logout")

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	delete(i.Sessions, session.Value)
}

func register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	utils.SetupCORS(&w)
	if r.Method == http.MethodOptions {
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
	storeUser(&newUser)
	fmt.Printf("New user. Private user data: %+v\n", newUser)

	responseBody := []byte("{\"userID\":" + newUser.ID + "}")
	w.Write(responseBody)
}

func getUser(user i.User) (i.User, error) {
	if u, ok := i.Users[user.Login]; ok {
		if utils.HashPassword(user.Password) == u.Password {
			return u, nil
		}
	}
	return user, errors.New("no user")
}

func storeUser(u *i.User) {
	i.IDCounter++
	u.ID = "id" + fmt.Sprint(i.IDCounter)
	u.Password = utils.HashPassword(u.Password)
	i.IDToLogin[u.ID] = u.Login
	i.Users[u.Login] = *u
}
