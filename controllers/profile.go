package controllers

import (
	"net/http"
	"vk.com/models"
)

func getUserBySession(user models.User)  {
	//Users find by field session
}

var inProfile string = `{"Auth":"success"}`

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