package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GenSession(id string) (session string) {
	hash := sha256.Sum256([]byte(id + fmt.Sprint(time.Now().Unix())))
	session = fmt.Sprintf("%x", hash)
	return
}

func Hash(password string) string {
	hash := sha256.Sum256([]byte(password + Salt))
	return fmt.Sprintf("%x", hash)
}

