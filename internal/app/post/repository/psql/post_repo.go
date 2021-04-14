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

type PostRepoPsql struct {
	DB *gorm.DB
}

func NewPostRepoPsql(db *gorm.DB) *PostRepoPsql {
	return &PostRepoPsql{DB: db}
}

func (urp *PostRepoPsql) SavePost(ctx context.Context, newPost *models.Post, userOwner *models.User, fileHandlers map[string][]*multipart.FileHeader) error {
	newPost.AuthorId = userOwner.ID
	err := urp.DB.Model(userOwner).Association("Posts").Append(newPost)
	//urp.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(userOwner)
	if err != nil {
		return err
	}
	urp.storeImg(ctx, newPost, fileHandlers)
	fmt.Println("Post added")
	return nil
}

func (urp *PostRepoPsql) storeImg(ctx context.Context, newPost *models.Post, fileHandlers map[string][]*multipart.FileHeader) error {

	for name, fileHeader := range fileHandlers {
		newImg := models.Img{}
		file, err := fileHeader[0].Open()
		if err != nil {
			return err
		}
		t := time.Now()
		salt := fmt.Sprintf(t.Format(time.RFC3339))
		genFileName := hasher.Hash(name + salt)
		defer file.Close()
		localImg, err := os.OpenFile("./static/posts/"+genFileName + ".png", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		newImg.Url = "/static/posts/" + genFileName + ".png"
		newImg.PostID = newPost.ID
		err = urp.DB.Create(&newImg).Error
		if err != nil {
			return err
		}
		defer localImg.Close()
		_, _ = io.Copy(localImg, file)
	}
	return nil
}

func (urp *PostRepoPsql) GetPosts(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	err := urp.DB.WithContext(ctx).Preload("UrlImgs").Find(&posts).Error
	if err != nil {
		return nil, err
	}
	for i, _ := range posts {
		u := models.User{}
		err = urp.DB.WithContext(ctx).First(&u, "id = ?", posts[i].AuthorId).Error
		if err != nil {
			return nil, err
		}
		posts[i].Author = u.FirstName + " " + u.LastName
		posts[i].AuthorAva = u.Avatar
	}
	return posts, nil
}

func (urp *PostRepoPsql) GetUserPosts(ctx context.Context, u *models.User) ([]models.Post, error) {
	var posts []models.Post
	err := urp.DB.WithContext(ctx).Preload("UrlImgs").Where("author_id = ?", u.ID).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (urp *PostRepoPsql) GetPost(ctx context.Context, id int) (*models.Post, error) {
	return nil, nil
}
