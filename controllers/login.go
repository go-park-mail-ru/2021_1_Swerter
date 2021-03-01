package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"vk.com/models"
)


var loginSuccess string = `{"status":"true"}`
var loginFail string = `{"status":"false"}`

var logoutSuccess string = `{"status":"true"}`

func isUserExist(user models.User) bool {
	for _, u := range Users {
		if  u.Login == user.Login && user.Password == u.Password{
			return true
		}
	}
	return false
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		return
	}

	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	decoder.Decode(&user)
	fmt.Printf("%+v\n", user)

	if isUserExist(user) {
		expiration := time.Now().Add(10 * time.Hour)
		cookie := http.Cookie{
			Name:    "session_id",
			Value:   user.Login,
			Expires: expiration,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		w.Write([]byte(loginSuccess))
	} else {
		w.Write([]byte(loginFail))
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	w.Write([]byte(logoutSuccess))
}