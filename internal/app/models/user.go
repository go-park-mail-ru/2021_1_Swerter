package models

import (
	"context"
	"mime/multipart"
)

//TODO: validate
type User struct {
	ID          string        `json:"userId"`
	Login       string        `json:"login"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	OldPassword string        `json:"oldPassword"`
	Password    string        `json:"password"`
	Posts       map[int]*Post `json:"postsData"`
	Avatar      string        `json:"avatar"`
}

type UserUsecase interface {
	SaveUser(ctx context.Context, u *User) error
	LoginUser(c context.Context, user *User) (*Session, error)
	GetPrivateUser(ctx context.Context, login string, password string) (*User, error)
	GetUserBySession(c context.Context, sessionValue string) (*User, error)
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, newUser *User, session string) error
	UploadAvatar(c context.Context, sessionId string, file multipart.File) error
	//Для логина и пароля
	UpdateSecureUser(ctx context.Context, userId int, login string, pass string, newUser *User) error
}

type UserRepository interface {
	SaveUser(ctx context.Context, u *User) error
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	GetPrivateUser(ctx context.Context, login string, password string) (*User, error)
	UpdateUser(ctx context.Context, oldUser *User, newUser *User) error
	UploadAvatar(c context.Context, user *User, file multipart.File) error
}
