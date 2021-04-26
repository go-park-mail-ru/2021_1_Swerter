module my-motivation

// +heroku goVersion go1.15
go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.8.1
	gitlab.com/Burunduck/user-service v0.0.0-20210425124523-8aff2e0dfa3e
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.6
)
