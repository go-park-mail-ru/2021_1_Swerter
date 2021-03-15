package posts

import (
	m "my-motivation/apps/middleware"
	"github.com/gorilla/mux"
)

func SetupRouterPosts(r *mux.Router) {
	r.HandleFunc("", m.CORS(allPosts)).Methods("GET", "OPTIONS")
	r.HandleFunc("/add", m.CORS(addPost)).Methods("POST", "OPTIONS")
}
