package utils

import (
	"fmt"
	i "my-motivation/internal"
	"net/http"
)

func SessionToUser(r *http.Request) *i.User {
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		fmt.Println("no cookie")
		return nil
	}

	if isSessionExist(session.Value) {
		userID := i.Sessions[session.Value]
		user := i.Users[i.IDToLogin[userID]]
		return &user
	}
	fmt.Println("no seesiom")
	return nil
}

func isSessionExist(session string) bool {
	if _, ok := i.Sessions[session]; ok {
		return true
	}
	return false
}