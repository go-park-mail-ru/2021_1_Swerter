package models

import (
	"context"
	"gorm.io/gorm"
	"mime/multipart"
)

type Post struct {
	gorm.Model
	ID          int      `gorm:"primaryKey;autoIncrement:true"`
	Author      string   `json:"postCreator"`
	AuthorAva   string   `json:"imgAvatar"`
	AuthorId    int      `json:"postCreatorId"`
	Text        string   `json:"textPost"`
	UrlImgs     []Img 	 `json:"imgContent" gorm:"foreignKey:PostID"`
	Date        string   `json:"date"`
	Liked       bool     `json:"liked" gorm:"-"`
	LikeCounter int      `json:"likeCounter" gorm:"-"`
}

type PostsUsecase interface {
	SavePost(ctx context.Context, session string, fileHandlers map[string][]*multipart.FileHeader, post *Post) error
	GetPosts(ctx context.Context, session string) ([]Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}

type PostsRepository interface {
	SavePost(ctx context.Context, post *Post, userOwner *User, fileHandlers map[string][]*multipart.FileHeader) error
	GetPosts(ctx context.Context) ([]Post, error)
	GetUserPosts(ctx context.Context, u *User) ([]Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}
