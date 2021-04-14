package psql

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"my-motivation/internal/app/models"
	"time"
)

type SessionsManagerPsql struct {
	DB *gorm.DB
}

func NewSessionsManagerPsql(db *gorm.DB) *SessionsManagerPsql {
	return &SessionsManagerPsql{DB: db}
}

func (sm *SessionsManagerPsql) Create(userID int) (*models.Session, error) {
	sessionId := genSession(userID)
	sess := models.Session{
		ID:     sessionId,
		UserID: userID,
	}
	err := sm.DB.Create(&sess).Error
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

func (sm *SessionsManagerPsql) GetUserId(sessionId string) (int, error) {
	sess := models.Session{}
	err := sm.DB.First(&sess, "id = ?", sessionId).Error
	if err != nil {
		return 0, errors.New("No session")
	}
	return sess.UserID, nil
}

func genSession(id int) (session string) {
	hash := sha256.Sum256([]byte(fmt.Sprint(id) + fmt.Sprint(time.Now().Unix())))
	session = fmt.Sprintf("%x", hash)
	return
}
