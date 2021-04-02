package models

import (
	"context"
)

//TODO: validate
type User struct {
	ID          string       `json:"userId"`
	Login       string       `json:"login"`
	FirstName   string       `json:"firstName"`
	LastName    string       `json:"lastName"`
	OldPassword string       `json:"oldPassword"`
	Password    string       `json:"password"`
	PostsId     []int 		 `json:"postsData"`
	Avatar      string       `json:"avatar"`
}

type UserUsecase interface {
	SaveUser(ctx context.Context, u *User) error
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	GetUserById(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, userId int, newUser *User) error
	//Для логина и пароля
	UpdateSecureUser(ctx context.Context, userId int, login string, pass string, newUser *User) error
}

type UserRepository interface {
	SaveUser(ctx context.Context, u *User) error
	GetUserByLogin(ctx context.Context, login string) (*User, error)
	UpdateUser(ctx context.Context, userId int, newUser *User) error
}

