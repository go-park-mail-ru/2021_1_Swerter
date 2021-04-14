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
	GetAlbum(c context.Context, session string, albumID int) (*Album, error)
}

type AlbumRepository interface {
	SaveAlbum(ctx context.Context, album *Album, userOwner *User, fileHandlers map[string][]*multipart.FileHeader) error
	GetUserAlbums(ctx context.Context, u *User) ([]Album, error)
	GetAlbum(ctx context.Context, albumID int) (*Album, error)
}
