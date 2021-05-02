package models

import (
	"context"
	"gorm.io/gorm"
	"mime/multipart"
)

type User struct {
	gorm.Model
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	Login       string  `json:"login"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	OldPassword string  `json:"oldPassword"`
	Password    string  `json:"password"`
	Avatar      string  `json:"avatar"`
	IsNotified  bool    `json:"isNotified"`
	IsFriend    bool    `json:"isFriend"`
	Posts       []Post  `json:"postsData" gorm:"foreignKey:AuthorId"`
	Albums      []Album `json:"albumsData" gorm:"foreignKey:AuthorId"`
}



func (u *User) Public() User {
	return User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Avatar:     u.Avatar,
		Posts:      u.Posts,
		IsNotified: u.IsNotified,
		IsFriend:   u.IsFriend,
	}
}

func (u *User) Private() User {
	return User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Avatar:    u.Avatar,
		Posts:     u.Posts,
		Login:     u.Login,
	}
}

type UserUsecase interface {
	Register(ctx context.Context, u *User) error
	Login(c context.Context, user *User) (string, error)
	Logout(ctx context.Context, session string) error

	GetPrivateUser(ctx context.Context, login string, password string) (*User, error)
	GetUserBySession(c context.Context, sessionValue string) (*User, error)
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, newUser *User, session string) error
	UploadAvatar(c context.Context, sessionId string, file multipart.File) error
	//Для логина и пароля
	UpdateSecureUser(ctx context.Context, userId int, login string, pass string, newUser *User) error
}

type UserRepository interface {
	SaveUser(ctx context.Context, u *User) error
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	GetPrivateUser(ctx context.Context, login string, password string) (*User, error)
	UpdateUser(ctx context.Context, oldUser *User, newUser *User) error
	UploadAvatar(c context.Context, user *User, file multipart.File) error
	SearchUsersByFullName(ctx context.Context, userName string, userSurname string) ([]User, error)
	SearchUsersByName(ctx context.Context, userName string) ([]User, error)
	SearchUsersBySurname(ctx context.Context, userSurname string) ([]User, error)
}
