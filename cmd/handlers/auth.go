package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	i "vk.com/internal"
)

var registerSuccess string = `"status":"true"`
var registerFail string = `"status":"false"`
var loginSuccess string = `{"status":"true"}`
var loginFail string = `{"status":"false"}`
var logoutSuccess string = `{"status":"true"}`
var inProfile string = `{"Auth":"success"}`

//var Auth handler = handler{}

//type handler struct {
//}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		return
	}

	decoder := json.NewDecoder(r.Body)
	newUser := i.User{}
	decoder.Decode(&newUser)
	i.IDCounter++
	newUser.ID = i.IDCounter

	i.Users[newUser.Login] = newUser

	fmt.Printf("%+v\n", newUser)
	w.Write([]byte(registerSuccess))
	http.Redirect(w, r, "/", http.StatusFound)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		return
	}

	decoder := json.NewDecoder(r.Body)
	user := i.User{}
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
		i.SessionsCounter++
		i.Sessions[i.SessionsCounter] = user
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

func ProfilePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		return
	}

	_, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", 401)
		return
	} else {
		w.Write([]byte(inProfile))
		return
	}

}

func isUserExist(user i.User) bool {
	if u, ok := i.Users[user.Login]; ok {
		if u.Password == user.Password {
			return true
		}
	}
	return false
}