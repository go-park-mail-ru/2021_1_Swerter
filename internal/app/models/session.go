package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	ID     string `gorm:"primaryKey;autoIncrement:false"`
	UserID int
}

type SessionRepository interface {
	CreateSession(userId int) (Session, error)
	UpdateSession()
	DeleteSession()
	GetSession()
}
