package session

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"my-motivation/internal/app/models"
	"sync"
	"time"
)

type SessionsManager struct {
	sessions map[string]*models.Session
	sessionCounter int
	mu   *sync.Mutex
}

func NewSessionManager() *SessionsManager {
	return &SessionsManager{
		sessions: make(map[string]*models.Session, 10),
		mu:   &sync.Mutex{},
	}
}

func (sm *SessionsManager) Create(userID string) (*models.Session, error) {
	sessionId := genSession(userID)
	sess := models.Session{
		ID:     sessionId,
		UserID: userID,
	}

	sm.mu.Lock()
	sm.sessions[sess.ID] = &sess
	sm.sessionCounter++
	sm.mu.Unlock()

	return &sess, nil
}

func (sm *SessionsManager) GetUserId(sessionValue string) (userId string, err error){
	session, ok := sm.sessions[sessionValue]
	if !ok || session == nil {
		return "", errors.New("no session")
	}
	return session.UserID, nil
}

func genSession(id string) (session string) {
	hash := sha256.Sum256([]byte(id + fmt.Sprint(time.Now().Unix())))
	session = fmt.Sprintf("%x", hash)
	return
}

