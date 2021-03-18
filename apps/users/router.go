package users

import (
	"github.com/gorilla/mux"
	m "my-motivation/apps/middleware"
)

func SetupRouterUsers(r *mux.Router) {
	r.HandleFunc("/loadImg", m.CORS(UploadFile)).Methods("POST")
	r.HandleFunc("", m.CORS(userProfile)).Methods("GET", "POST", "OPTIONS")
	r.HandleFunc("/{userID}", m.CORS(getUserProfileByID)).Methods("GET", "OPTIONS")
}
