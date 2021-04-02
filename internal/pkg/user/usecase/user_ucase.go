package usecase

import (
	"context"
	"my-motivation/internal/pkg/models"
	"time"
)

type UserUsecase struct {
	userRepo models.UserRepository
	postRepo models.PostsRepository
	contextTimeout time.Duration
}

func NewUserUsecase(u models.UserRepository, p models.PostsRepository, t time.Duration) models.UserUsecase {
	return &UserUsecase{
		userRepo: u,
		postRepo: p,
		contextTimeout : t,
	}
}

func (uu *UserUsecase) SaveUser(ctx context.Context, u *models.User) error {
	return nil
}

func (uu *UserUsecase) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	return nil, nil
}

func (uu *UserUsecase) GetUserById(ctx context.Context, id int) (*models.User, error) {
	return nil,	nil
}

func (uu *UserUsecase) UpdateUser(ctx context.Context, userId int, newUser *models.User) error {
	return nil
}

//Для логина и пароля
func (uu * UserUsecase) UpdateSecureUser(ctx context.Context, userId int, login string, pass string, newUser *models.User) error {
	return nil
}

