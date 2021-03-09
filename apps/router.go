package apps

import (
	"my-motivation/apps/auth"
	"my-motivation/apps/news"
	"my-motivation/apps/users"

	"github.com/gorilla/mux"
)

func SetupRouterMain(mainRouter *mux.Router) {
	authRouter := mainRouter.PathPrefix("/").Subrouter()
	auth.SetupRouterAuth(authRouter)

	usersRouter := mainRouter.PathPrefix("/users").Subrouter()
	users.SetupRouterUsers(usersRouter)

	postsRouter := mainRouter.PathPrefix("/posts").Subrouter()
	news.SetupRouterPosts(postsRouter)
}
