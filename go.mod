module my-motivation

// +heroku goVersion go1.15
go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.6
)
