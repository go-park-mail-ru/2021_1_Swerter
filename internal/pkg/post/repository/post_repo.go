package repository

import (
	"context"
	"my-motivation/internal/pkg/models"
)

type PostRepo struct {
	data map[string]*models.Post
}

func NewPostRepo() models.PostsRepository {
	return &PostRepo{
		data: map[string]*models.Post{},
	}
}

func (ur * PostRepo) SavePost(ctx context.Context, u *models.Post) error {
	return nil
}

func (ur * PostRepo) GetPosts(ctx context.Context) ([]models.Post, error) {
	return nil, nil
}

func (ur * PostRepo) GetPost(ctx context.Context, id int) (* models.Post, error) {
	return nil, nil
}
