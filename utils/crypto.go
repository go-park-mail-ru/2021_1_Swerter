package utils

import (
	"crypto/sha256"
	"fmt"
	i "my-motivation/internal"
	"time"
)

func GenSession(id string) (session string) {
	hash := sha256.Sum256([]byte(id + fmt.Sprint(time.Now().Unix())))
	session = fmt.Sprintf("%x", hash)
	return
}

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password + i.Salt))
	return fmt.Sprintf("%x", hash)
}
