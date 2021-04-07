package main

//Запускать go run в этой директории, т.к. go run делает бинарник, а пути у нас до staticfileHandlera захардкожены

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"my-motivation/internal/app/middleware"
	"my-motivation/internal/app/models"
	"my-motivation/internal/app/session"
	"my-motivation/internal/pkg/utils/logger"

	_postHttpDelivery "my-motivation/internal/app/post/delivery/http"
	_postUsecase "my-motivation/internal/app/post/usecase"
	_userHttpDelivery "my-motivation/internal/app/user/delivery/http"
	_userUsecase "my-motivation/internal/app/user/usecase"
	//_postRepo "my-motivation/internal/app/post/repository"
	_postRepo "my-motivation/internal/app/post/repository/psql"
	//_userRepo "my-motivation/internal/app/user/repository"
	_userRepoPsql "my-motivation/internal/app/user/repository/psql"

	"net/http"
	"os"
	"time"
)


func getPostgres() *gorm.DB {
	dsn := "host=localhost user=vk password=vk dbname=vk port=5400 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//Только во ремя разработки автомигрете
	db.AutoMigrate(&models.User{}, &models.Post{})
	return db
}


//
type Topic struct {
	gorm.Model
	// TopicID uint `gorm:"primary_key"`
	Name    string
	Posts   []Post `gorm:"ForeignKey:ID"`
}

type Post struct {
	gorm.Model
	// PostID     uint `gorm:"primary_key"`
	Title      string
	TopicRefer uint `gorm:"column:id"`
}
//

func main() {
	//logger
	log := logger.NewLogger()

	//repo
	userRepo := _userRepoPsql.NewUserRepoPsql(getPostgres())
	postRepo := _postRepo.NewPostRepoPsql(getPostgres())
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
