package psql

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/hasher"
	"os"
	"time"
)

type AlbumRepoPsql struct {
	DB *gorm.DB
}

func NewAlbumRepoPsql(db *gorm.DB) *AlbumRepoPsql {
	return &AlbumRepoPsql{DB: db}
}

func (arp *AlbumRepoPsql) SaveAlbum(ctx context.Context, newAlbum *models.Album, userOwner *models.User, fileHandlers map[string][]*multipart.FileHeader) error {
	newAlbum.AuthorId = userOwner.ID
	err := arp.DB.Model(userOwner).Association("Albums").Append(newAlbum)
	//urp.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(userOwner)
	if err != nil {
		return err
	}
	err = arp.storeImg(ctx, newAlbum, fileHandlers)
	fmt.Println(err)
	fmt.Println("Album added")
	return nil
}

func (arp *AlbumRepoPsql) storeImg(ctx context.Context, newAlbum *models.Album, fileHandlers map[string][]*multipart.FileHeader) error {
	for name, fileHeader := range fileHandlers {
		newImg := models.AlbumImg{}
		file, err := fileHeader[0].Open()
		if err != nil {
			return err
		}
		t := time.Now()
		salt := fmt.Sprintf(t.Format(time.RFC3339))
		genFileName := hasher.Hash(name + salt)
		defer file.Close()
		localImg, err := os.OpenFile("../../static/albums/"+ genFileName + ".png", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		newImg.Url = "/static/albums/" + genFileName + ".png"
		newImg.AlbumID = newAlbum.ID
		err = arp.DB.Create(&newImg).Error
		if err != nil {
			return err
		}
		defer localImg.Close()
		_, _ = io.Copy(localImg, file)
	}
	return nil
}