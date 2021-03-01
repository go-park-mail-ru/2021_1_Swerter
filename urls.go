package main

import (
	"net/http"
	c "vk.com/controllers"
)

func routes(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/register", c.RegisterPage)
	mux.HandleFunc("/login", c.LoginPage)
	mux.HandleFunc("/logout", c.LogoutPage)
	mux.HandleFunc("/profile", c.ProfilePage)
}

