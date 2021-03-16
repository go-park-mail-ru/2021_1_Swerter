package internal

import (
	"crypto/sha256"
	"fmt"
)

var IDToLogin = make(map[string]string)

//Login: User
var Users = map[string]User{}
var IDCounter int

type User struct {
	ID          string       `json:"userId"`
	Login       string       `json:"login"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	OldPassword string       `json:"oldPassword"`
	Password    string       `json:"password"`
	Posts       map[int]Post `json:"postsData"`
	Avatar      string       `json:"avatar"`
}

func UpdateUser(newUser *User, oldUser *User) {
	newUser.ID = oldUser.ID

	if newUser.Login == "" {
		newUser.Login = oldUser.Login
	} else {
		IDToLogin[newUser.ID] = newUser.Login
	}

	if newUser.FirstName == "" {
		newUser.FirstName = oldUser.FirstName
	}

	if newUser.LastName == "" {
		newUser.LastName = oldUser.LastName
	}

	newUser.Posts = oldUser.Posts
	newUser.Avatar = oldUser.Avatar
	delete(Users, oldUser.Login)
	Users[newUser.Login] = *newUser
}

func HashPassword(password string) string {
	var salt = "super_secure_key"
	hash := sha256.Sum256([]byte(password + salt))
	return fmt.Sprintf("%x", hash)
}
