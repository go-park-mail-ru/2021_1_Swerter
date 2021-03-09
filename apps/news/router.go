package news

import (
	"github.com/gorilla/mux"
)

func SetupRouterPosts(r *mux.Router) {
	r.HandleFunc("", posts)
}
