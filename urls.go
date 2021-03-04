package main

import (
	"net/http"
	h "vk.com/cmd/handlers"
)

func routes(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/register", h.RegisterPage)
	mux.HandleFunc("/login", h.LoginPage)
	mux.HandleFunc("/logout", h.LogoutPage)
	mux.HandleFunc("/profile", h.ProfilePage)
}

