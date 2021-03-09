package apps

import (
	"./auth"
	"./users"
	"github.com/gorilla/mux"
)

func SetupRouterMain(mainRouter *mux.Router) {
	authRouter := mainRouter.PathPrefix("/").Subrouter()
	auth.SetupRouterAuth(authRouter)

	usersRouter := mainRouter.PathPrefix("/profile").Subrouter()
	users.SetupRouterUsers(usersRouter)
}
