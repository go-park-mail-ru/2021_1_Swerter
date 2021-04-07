package models

import (
	"context"
	"gorm.io/gorm"
	"mime/multipart"
)

//TODO: validate
type Post struct {
	gorm.Model
	ID        int    `gorm:"primaryKey;autoIncrement:true"`
	Author    string `json:"postCreator"` //устанавливаю при подзгрузке на фронте, если отдаю в профиль
	AuthorAva string `json:"imgAvatar"`   //устанавливаю при подзгрузке на фронте, если отдаю в профиль
	AuthorId  int    `json:"postCreatorId"`
	Text      string `json:"textPost"`
	UrlImg    string `json:"imgContent"`
	Date      string `json:"date"`
}

type PostsUsecase interface {
	SavePost(ctx context.Context, session string, imgFile multipart.File, fileHandler *multipart.FileHeader, post *Post) error
	GetPosts(ctx context.Context, session string) ([]Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}

type PostsRepository interface {
	SavePost(ctx context.Context, post *Post, userOwner *User, file multipart.File, fileHandler *multipart.FileHeader) error
	GetPosts(ctx context.Context) ([]Post, error)
	GetUserPosts(ctx context.Context, u *User) ([]Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}
