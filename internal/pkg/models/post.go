package models

import "context"

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
	SavePost(ctx context.Context, u *Post) error
	GetPosts(ctx context.Context) ([]Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}

type PostsRepository interface {
	SavePost(ctx context.Context, u *Post) error
	GetPosts(ctx context.Context) ([]Post, error)
	GetPost(ctx context.Context, id int) (*Post, error)
}