package hasher

import (
	"crypto/sha256"
	"fmt"
)

func Hash(password string) string {
	hash := sha256.Sum256([]byte(password + Salt))
	return fmt.Sprintf("%x", hash)
}

