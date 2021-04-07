package psql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/hasher"
	"os"
	"time"
)

type PostRepoPsql struct {
	DB *gorm.DB
}

func NewPostRepoPsql(db *gorm.DB) *PostRepoPsql {
	return &PostRepoPsql{DB: db}
}

func (urp *PostRepoPsql) SavePost(ctx context.Context, newPost *models.Post, userOwner *models.User, file multipart.File, fileHandler *multipart.FileHeader) error {
	newPost.AuthorId = userOwner.ID
	urp.storeImg(ctx, newPost, file, fileHandler)
	err := urp.DB.Model(userOwner).Association("Posts").Append(newPost)
	//urp.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(userOwner)
	if err != nil {
		return err
	}
	fmt.Println("Post added")
	return nil
}

func (urp *PostRepoPsql) storeImg(ctx context.Context, newPost *models.Post, file multipart.File, fileHandler *multipart.FileHeader) error {
	if file == nil || fileHandler.Filename == "" {
		return errors.New("empty img file")
	}
	t := time.Now()
	salt := fmt.Sprintf(t.Format(time.RFC3339))
	genFileName := hasher.Hash(fileHandler.Filename + salt)

	defer file.Close()
	localImg, err := os.OpenFile("../../static/posts/"+genFileName, os.O_WRONLY|os.O_CREATE, 0666)
	newPost.UrlImg = "/static/posts/" + genFileName
	if err != nil {
		return err
	}

	defer localImg.Close()
	_, _ = io.Copy(localImg, file)
	return nil
}

func (urp *PostRepoPsql) GetPosts(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	err := urp.DB.WithContext(ctx).Find(&posts).Error
	if err != nil {
		return nil,err
	}
	return posts, nil
}

func (urp *PostRepoPsql) GetUserPosts(ctx context.Context, u *models.User) ([]models.Post, error) {
	var posts []models.Post
	err := urp.DB.WithContext(ctx).Where("author_id = ?", u.ID).Find(&posts).Error
	if err != nil {
		return nil,err
	}
	return posts, nil
}


func (urp *PostRepoPsql) GetPost(ctx context.Context, id int) (*models.Post, error) {
	return nil, nil
}


