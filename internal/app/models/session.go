package models

type Session struct {
	ID     string
	UserID string
}

type SessionRepository interface {
	CreateSession(userId string) (Session, error)
	UpdateSession()
	DeleteSession()
	GetSession()

}
