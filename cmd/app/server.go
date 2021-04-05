package main

import (
	"github.com/gorilla/mux"
	"my-motivation/internal/app/middleware"
	_postHttpDelivery "my-motivation/internal/app/post/delivery/http"
	_postRepo "my-motivation/internal/app/post/repository"
	_postUsecase "my-motivation/internal/app/post/usecase"
	"my-motivation/internal/app/session"
	_userHttpDelivery "my-motivation/internal/app/user/delivery/http"
	_userRepo "my-motivation/internal/app/user/repository"
	_userUsecase "my-motivation/internal/app/user/usecase"
	"net/http"
	"os"
	"time"
)

func main() {
	//repo
	userRepo := _userRepo.NewUserRepo()
	postRepo := _postRepo.NewPostRepo(userRepo)
	sessionManager := session.NewSessionManager()

	//usecase
	timeoutContext := 5 * time.Second
	userUsecase := _userUsecase.NewUserUsecase(userRepo, postRepo, timeoutContext, sessionManager)
	postUsecase := _postUsecase.NewPostUsecase(userRepo, postRepo, timeoutContext, sessionManager)

	//delivery
	r := mux.NewRouter()
	_userHttpDelivery.NewUserHandler(r, userUsecase)
	_postHttpDelivery.NewPostHandler(r, postUsecase)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/"))))

	//index
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Kumusta Higala"))
	})

	//middleware
	handler := middleware.CORS(r)

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: handler,
	}

	server.ListenAndServe()
}
