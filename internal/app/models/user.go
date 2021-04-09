package models

import (
	"context"
	"gorm.io/gorm"
	"mime/multipart"
)

type User struct {
	gorm.Model
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	Login       string `json:"login"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
	Posts       []Post `json:"postsData" gorm:"foreignKey:AuthorId"`
	Avatar      string `json:"avatar"`
}

func (u *User) Public() User {
	return User{
		ID: u.ID,
		FirstName: u.FirstName,
		LastName: u.LastName,
		Avatar: u.Avatar,
		Posts: u.Posts,
	}
}

func (u *User) Private() User {
	return User{
		ID: u.ID,
		FirstName: u.FirstName,
		LastName: u.LastName,
		Avatar: u.Avatar,
		Posts: u.Posts,
		Login: u.Login,
	}
}

type UserUsecase interface {
	SaveUser(ctx context.Context, u *User) error
	LoginUser(c context.Context, user *User) (*Session, error)
	GetPrivateUser(ctx context.Context, login string, password string) (*User, error)
	GetUserBySession(c context.Context, sessionValue string) (*User, error)
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, newUser *User, session string) error
	UploadAvatar(c context.Context, sessionId string, file multipart.File) error
	AddFriend(c context.Context, session string, userFiend *User) error
	GetFriends(c context.Context, session string) (map[string]*User, error)
	//Для логина и пароля
	UpdateSecureUser(ctx context.Context, userId int, login string, pass string, newUser *User) error
}

type UserRepository interface {
	SaveUser(ctx context.Context, u *User) error
	SaveFriend(ctx context.Context, user *User, userFiend *User) error
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	GetPrivateUser(ctx context.Context, login string, password string) (*User, error)
	UpdateUser(ctx context.Context, oldUser *User, newUser *User) error
	GetFriends(ctx context.Context, user *User) (map[string]*User, error)
	UploadAvatar(c context.Context, user *User, file multipart.File) error
}
