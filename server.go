package main

import (
	_postRepo "my-motivation/internal/pkg/post/repository"
	_postUsecase "my-motivation/internal/pkg/post/usecase"
	_userRepo "my-motivation/internal/pkg/user/repository"
	_userUsecase "my-motivation/internal/pkg/user/usecase"
	"net/http"
	"os"
	"time"

	"my-motivation/apps"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	apps.SetupRouterMain(mainRouter)


	postRepo := _postRepo.NewPostRepo()
	userRepo := _userRepo.NewUserRepo()

	timeoutContext := 5 * time.Second
	userUsecase := _userUsecase.NewUserUsecase(userRepo, postRepo, timeoutContext)
	postUsercase := _postUsecase.NewPostUsecase(userRepo, postRepo, timeoutContext)

	//Кладем юскейсы в деливири

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: mainRouter,
	}

	server.ListenAndServe()
}
