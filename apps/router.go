package apps

import (
	"github.com/gorilla/mux"
	"my-motivation/apps/auth"
	"my-motivation/apps/news"
	"my-motivation/apps/users"
)

func SetupRouterMain(mainRouter *mux.Router) {
	authRouter := mainRouter.PathPrefix("/").Subrouter()
	auth.SetupRouterAuth(authRouter)

	usersRouter := mainRouter.PathPrefix("/profile").Subrouter()
	users.SetupRouterUsers(usersRouter)

	postsRouter := mainRouter.PathPrefix("/posts").Subrouter()
	news.SetupRouterPosts(postsRouter)

}
