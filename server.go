package main

import (
	"net/http"

	"./apps"
	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	apps.SetupRouterMain(mainRouter)

	server := http.Server{
		Addr:    ":8000",
		Handler: mainRouter,
	}

	server.ListenAndServe()
}
