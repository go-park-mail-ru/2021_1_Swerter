package models

import (
	"context"
	"mime/multipart"
)

//TODO: validate
type Post struct {
	Id        int
	Author    string `json:"postCreator"`//устанавливаю при подзгрузке
	AuthorAva string `json:"imgAvatar"` //устанавливаю при подзгрузке
	AuthorId  string `json:"postCreatorId"`
	Text      string `json:"textPost"`
	UrlImg    string `json:"imgContent"`
	Date      string `json:"date"`
}

type PostsUsecase interface {
	SavePost(ctx context.Context, session string, imgFile multipart.File, fileHandler *multipart.FileHeader, post *Post) error
	GetPosts(ctx context.Context, session string) (map[int]*Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}

type PostsRepository interface {
	SavePost(ctx context.Context, post *Post, userOwner *User, file multipart.File, fileHandler *multipart.FileHeader) error
	GetPosts(ctx context.Context) (map[int]*Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}