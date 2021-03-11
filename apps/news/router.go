package news

import (
	"github.com/gorilla/mux"
)

func SetupRouterPosts(r *mux.Router) {
	r.HandleFunc("", allPosts).Methods("GET", "OPTIONS")
	r.HandleFunc("/add", addPost).Methods("POST", "OPTIONS")
}
