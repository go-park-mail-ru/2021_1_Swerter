package main

import (
	"net/http"
	c "vk.com/controllers"
)

func routes(mux *http.ServeMux) {
	mux.HandleFunc("/login", c.LoginPage)
	mux.HandleFunc("/logout", c.LogoutPage)
	mux.Handle("/", http.FileServer(http.Dir("./static")))
}

