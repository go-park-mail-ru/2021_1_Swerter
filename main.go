package main

import (
	"fmt"
	"net/http"
)

func runServer(addr string) {
	mux := http.NewServeMux()
	routes(mux)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	fmt.Println("starting server at", addr)
	server.ListenAndServe()
}

func main() {
	runServer(":3000")
}
