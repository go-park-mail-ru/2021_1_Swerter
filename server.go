package main

import (
	"net/http"
	"os"

	"my-motivation/apps"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	apps.SetupRouterMain(mainRouter)

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mainRouter,
	}

	server.ListenAndServe()
}
