package utils

import (
	i "my-motivation/internal"
	"net/http"
)

func SessionToUser(r *http.Request) *i.User {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		return nil
	}

	if isSessionExist(session.Value) {
		userID := i.Sessions[session.Value]
		user := i.Users[i.IDToLogin[userID]]
		return &user
	}

	return nil
}

func isSessionExist(session string) bool {
	if _, ok := i.Sessions[session]; ok {
		return true
	}
	return false
}