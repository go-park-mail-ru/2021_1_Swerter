package main

import (
	"net/http"

	"my-motivation/apps"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	apps.SetupRouterMain(mainRouter)

	server := http.Server{
		Addr:    ":80",
		Handler: mainRouter,
	}

	server.ListenAndServe()
}
