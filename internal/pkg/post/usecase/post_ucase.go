package usecase

import (
	"context"
	"my-motivation/internal/pkg/models"
	"time"
)

type PostUsecase struct {
	UserRepo models.UserRepository
	PostRepo models.PostsRepository
	contextTimeout time.Duration
}

func NewPostUsecase(ur models.UserRepository, pr models.PostsRepository, timeout time.Duration) models.PostsUsecase {
	return &PostUsecase{
		UserRepo: ur,
		PostRepo: pr,
		contextTimeout: timeout,
	}
}

func (pu * PostUsecase) SavePost(ctx context.Context, u *models.Post) error {
	return nil
}

func (pu * PostUsecase) GetPosts(ctx context.Context) ([]models.Post, error) {
	return nil, nil
}

func (pu * PostUsecase) GetPost(ctx context.Context, id int) (* models.Post, error) {
	return nil, nil
}

