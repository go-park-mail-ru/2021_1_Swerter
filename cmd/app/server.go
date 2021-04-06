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
	"my-motivation/internal/pkg/utils/logger"

	//"my-motivation/internal/pkg/utils/logger"
	"net/http"
	"os"
	"time"
	//log "github.com/sirupsen/logrus"

)

func main() {
	//logger
	log := logger.NewLogger()

	//repo
	userRepo := _userRepo.NewUserRepo()
	postRepo := _postRepo.NewPostRepo(userRepo)
	sessionManager := session.NewSessionManager()

	//usecase
	timeoutContext := 2 * time.Second
	userUsecase := _userUsecase.NewUserUsecase(userRepo, postRepo, timeoutContext, sessionManager)
	postUsecase := _postUsecase.NewPostUsecase(userRepo, postRepo, timeoutContext, sessionManager)

	//delivery
	r := mux.NewRouter()
	_userHttpDelivery.NewUserHandler(r, userUsecase, log)
	_postHttpDelivery.NewPostHandler(r, postUsecase, log)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/"))))

	//index
	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Kumusta Higala"))
	})

	//middleware
	handler := middleware.CORS(r)
	handler = middleware.LoggingMiddleware(handler, log)

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: handler,
	}

	log.Log.Infof("Server start at %s port", server.Addr)
	server.ListenAndServe()
}
