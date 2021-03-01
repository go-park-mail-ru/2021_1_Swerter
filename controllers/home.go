package controllers

import (
	"fmt"
	"net/http"
	v "vk.com/view"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	loggedIn := (err != http.ErrNoCookie)

	if loggedIn {
		fmt.Fprintln(w,v.LogoutFormTmpl + session.Value)
	} else {
		fmt.Fprintln(w,v.LoginFormTmpl)
	}
}