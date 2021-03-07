package index

import "github.com/gorilla/mux"

func SetupRouterIndex(r *mux.Router) {
	r.HandleFunc("/", index)
}
