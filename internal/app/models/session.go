package models

type Session struct {
	ID     string
	UserID int
}

type SessionRepository interface {
	CreateSession(userId int) (Session, error)
	UpdateSession()
	DeleteSession()
	GetSession()

}
