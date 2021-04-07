package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"my-motivation/internal/app/middleware"
	"my-motivation/internal/app/models"

	//Драйвер к постгресу
	_postHttpDelivery "my-motivation/internal/app/post/delivery/http"
	_postRepo "my-motivation/internal/app/post/repository"
	_postUsecase "my-motivation/internal/app/post/usecase"
	"my-motivation/internal/app/session"
	_userHttpDelivery "my-motivation/internal/app/user/delivery/http"
	_userUsecase "my-motivation/internal/app/user/usecase"
	"my-motivation/internal/pkg/utils/logger"
	//_userRepo "my-motivation/internal/app/user/repository"
	_userRepoPsql "my-motivation/internal/app/user/repository/psql"

	"net/http"
	"os"
	"time"
	//log "github.com/sirupsen/logrus"

)

//TODO:вынести в конфиг всё
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func getPostgres() *gorm.DB {
	dsn := "host=localhost user=vk password=vk dbname=vk port=5400 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//Только во ремя разработки автомигрете
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&models.User{})

	db.Create(&Product{Code: "10", Price: 100})
	return db
}

func main() {
	//logger
	log := logger.NewLogger()

	//repo
	//userRepo := _userRepo.NewUserRepo()
	//NewUserRepoPsql
	userRepo := _userRepoPsql.NewUserRepoPsql(getPostgres())
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
