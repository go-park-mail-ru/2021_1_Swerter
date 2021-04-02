package repository

import (
	"my-motivation/internal/pkg/models"
	"my-motivation/utils"
	"sync"
)

type SessionsManager struct {
	data map[string]*models.Session
	mu   *sync.Mutex
}

func NewSessionManager() *SessionsManager {
	return &SessionsManager{
		data: make(map[string]*models.Session, 10),
		mu:   &sync.Mutex{},
	}
}

func (sm *SessionsManager) Create(userID string) (*models.Session, error) {
	sessionId := utils.GenSession(userID)
	sess := models.Session{
		ID:     sessionId,
		UserID: userID,
	}

	sm.mu.Lock()
	sm.data[sess.ID] = &sess
	sm.mu.Unlock()

	return &sess, nil
}
