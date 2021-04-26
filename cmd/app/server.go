package main

//Запускать go run в этой директории, т.к. go run делает бинарник, а пути у нас до staticfileHandlera захардкожены

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	_albumDelivery "my-motivation/internal/app/album/delivery/http"
	_albumRepo "my-motivation/internal/app/album/repository/psql"
	_albumUsecase "my-motivation/internal/app/album/usecase"
	_friendHttpDelivery "my-motivation/internal/app/friend/delivery/http"
	_friendRepo "my-motivation/internal/app/friend/repository/psql"
	_friendUsecase "my-motivation/internal/app/friend/usecase"
	"my-motivation/internal/app/middleware"
	"my-motivation/internal/app/models"
	user_service_client "my-motivation/internal/pkg/client/user-service-client"
	"my-motivation/internal/pkg/utils/logger"

	_likeHttpDelivery "my-motivation/internal/app/like/delivery/http"
	_likeRepoPsql "my-motivation/internal/app/like/repository/psql"
	_likeUsecase "my-motivation/internal/app/like/usecase"
	_postHttpDelivery "my-motivation/internal/app/post/delivery/http"
	_postRepo "my-motivation/internal/app/post/repository/psql"
	_postUsecase "my-motivation/internal/app/post/usecase"
	_userHttpDelivery "my-motivation/internal/app/user/delivery/http"
	_userRepoPsql "my-motivation/internal/app/user/repository/psql"
	_userUsecase "my-motivation/internal/app/user/usecase"
	//_userRepo "my-motivation/internal/app/user/repository"
	//_postRepo "my-motivation/internal/app/post/repository"
	_sessionManager "my-motivation/internal/app/session/psql"
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
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Session{}, &models.Friend{}, &models.Like{}, models.Img{}, models.Album{}, models.AlbumImg{})
	return db
}


func main() {
	//logger
	log := logger.NewLogger()

	//repo
	userRepo := _userRepoPsql.NewUserRepoPsql(getPostgres())
	postRepo := _postRepo.NewPostRepoPsql(getPostgres())
	friendRepo := _friendRepo.NewFriendRepoPsql(getPostgres())
	sessionManager := _sessionManager.NewSessionsManagerPsql(getPostgres())
	likeRepo := _likeRepoPsql.NewLikeRepoPsql(getPostgres())
	albumRepo := _albumRepo.NewAlbumRepoPsql(getPostgres())

	//services
	us, _ := user_service_client.New()

	//usecase
	timeoutContext := 2 * time.Second
	userUsecase := _userUsecase.NewUserUsecase(us, userRepo, postRepo, albumRepo, timeoutContext, sessionManager, likeRepo)
	postUsecase := _postUsecase.NewPostUsecase(userRepo, postRepo, timeoutContext, sessionManager, likeRepo)
	friendUsecase := _friendUsecase.NewFriendUsecase(friendRepo, userRepo, timeoutContext, sessionManager)
	likeUsecase := _likeUsecase.NewLikeUsecase(likeRepo, timeoutContext, sessionManager)
	albumUsecase := _albumUsecase.NewAlbumUsecase(userRepo, albumRepo, timeoutContext, sessionManager)

	//delivery
	r := mux.NewRouter()
	_userHttpDelivery.NewUserHandler(r, userUsecase, log)
	_postHttpDelivery.NewPostHandler(r, postUsecase, log)
	_friendHttpDelivery.NewFiendHandler(r, friendUsecase, log)
	_likeHttpDelivery.NewLikeHandler(r, likeUsecase, log)
	_albumDelivery.NewAlbumHandler(r, albumUsecase, log)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

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
