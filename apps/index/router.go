package index

import (
	"github.com/gorilla/mux"
	m "my-motivation/apps/middleware"
)

func SetupRouterIndex(r *mux.Router) {
	r.HandleFunc("/", m.CORS(index))
}
