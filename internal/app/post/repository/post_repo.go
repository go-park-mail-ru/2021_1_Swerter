package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/hasher"
	"os"
	"time"
)
//TODO: поправить везде инкапсуляцию
type PostRepo struct {
	//TODO:При переходе к бд решить проблему
	UserRepo    models.UserRepository
	Posts       map[int]*models.Post
	PostCounter int
}

func NewPostRepo(ur models.UserRepository) models.PostsRepository {
	return &PostRepo{
		Posts: map[int]*models.Post{},
		UserRepo: ur,
		PostCounter: 0,
	}
}

func (ur *PostRepo) SavePost(ctx context.Context, newPost *models.Post, userOwner *models.User, file multipart.File, fileHandler *multipart.FileHeader) error {
	//ur.PostCounter++
	//newPost.ID = ur.PostCounter
	//newPost.AuthorId = userOwner.ID
	//ur.storeImg(ctx, newPost, file, fileHandler)
	//ur.Posts[ur.PostCounter] = newPost
	//
	//userOwner.Posts[newPost.ID] = newPost
	return nil
}

func (ur *PostRepo) storeImg(ctx context.Context, newPost *models.Post, file multipart.File, fileHandler *multipart.FileHeader) error {
	if file == nil || fileHandler.Filename == "" {
		return errors.New("empty img file")
	}
	t := time.Now()
	salt := fmt.Sprintf(t.Format(time.RFC3339))
	genFileName := hasher.Hash(fileHandler.Filename + salt)

	defer file.Close()
	localImg, err := os.OpenFile("../../static/posts/"+genFileName, os.O_WRONLY|os.O_CREATE, 0666)
	newPost.UrlImgs = "/static/posts/" + genFileName
	if err != nil {
		fmt.Printf("Cant create file\n")
		return err
	}

	defer localImg.Close()
	_, _ = io.Copy(localImg, file)
	return nil
}

func (ur *PostRepo) GetPosts(ctx context.Context) (map[int]*models.Post, error) {
	curPosts := make(map[int]*models.Post)
	for k, v := range ur.Posts {
		u, err:= ur.UserRepo.GetUserById(ctx, v.AuthorId)
		if err != nil {
			return nil, err
		}
		v.Author = u.FirstName + " " + u.LastName
		v.AuthorAva = u.Avatar
		curPosts[k] = v
	}
	return curPosts, nil
}

func (ur *PostRepo) GetPost(ctx context.Context, id int) (*models.Post, error) {
	return nil, nil
}
