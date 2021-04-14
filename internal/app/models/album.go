package models

import (
	"context"
	"gorm.io/gorm"
	"mime/multipart"
)

//TODO: validate
type Album struct {
	gorm.Model
	ID          int        `gorm:"primaryKey;autoIncrement:true"`
	AuthorId    int        `json:"albumCreatorId"`
	Title       string     `json:"albumTitle"`
	Description string     `json:"albumDescription"`
	UrlImgs     []AlbumImg `json:"imgContent" gorm:"foreignKey:AlbumID"`
}

type AlbumsUsecase interface {
	SaveAlbum(ctx context.Context, session string, fileHandlers map[string][]*multipart.FileHeader, album *Album) error
	//GetPosts(ctx context.Context, session string) ([]Post, error)
	//GetPost(ctx context.Context, id int) (*Post, error)
}

type AlbumRepository interface {
	SaveAlbum(ctx context.Context, album *Album, userOwner *User, fileHandlers map[string][]*multipart.FileHeader) error
	//GetPosts(ctx context.Context) ([]Post, error)
	//GetUserPosts(ctx context.Context, u *User) ([]Post, error)
	//GetPost(ctx context.Context, id int) (*Post, error)
}
