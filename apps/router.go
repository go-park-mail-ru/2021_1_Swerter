package apps

import (
	"github.com/gorilla/mux"
	"my-motivation/apps/auth"
	"my-motivation/apps/posts"
	"my-motivation/apps/users"
	"net/http"
)

func SetupRouterMain(mainRouter *mux.Router) {
	authRouter := mainRouter.PathPrefix("/").Subrouter()
	auth.SetupRouterAuth(authRouter)

	usersRouter := mainRouter.PathPrefix("/profile").Subrouter()
	users.SetupRouterUsers(usersRouter)

	postsRouter := mainRouter.PathPrefix("/posts").Subrouter()
	posts.SetupRouterPosts(postsRouter)

	mainRouter.PathPrefix("/static/" ).Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}
