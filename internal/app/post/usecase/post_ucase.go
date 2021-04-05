package usecase

import (
	"context"
	"log"
	"mime/multipart"
	"my-motivation/internal/app/models"
	"my-motivation/internal/app/session"
	"time"
)

type PostUsecase struct {
	UserRepo       models.UserRepository
	PostRepo       models.PostsRepository
	contextTimeout time.Duration
	sessionManager *session.SessionsManager
}

func NewPostUsecase(ur models.UserRepository, pr models.PostsRepository, timeout time.Duration, sm *session.SessionsManager) models.PostsUsecase {
	return &PostUsecase{
		UserRepo: ur,
		PostRepo: pr,
		contextTimeout: timeout,
		sessionManager: sm,
	}
}

func (pu *PostUsecase) SavePost(c context.Context, session string, imgFile multipart.File, fileHandler *multipart.FileHeader, post *models.Post) error {
	userId, err := pu.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	userOwner, err :=pu.UserRepo.GetUserById(ctx, userId)
	if err != nil || userOwner == nil {
		log.Println("Add post failed")
		return err
	}

	err = pu.PostRepo.SavePost(ctx, post, userOwner, imgFile, fileHandler)
	if err != nil {
		return err
	}
	return nil
}

func (pu *PostUsecase) GetPosts(c context.Context, session string) (map[int]*models.Post, error) {
	userId, err := pu.sessionManager.GetUserId(session)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	//TODO: доступ к страницам для авторизованных пользователей вынести в мидлвару
	_, err =pu.UserRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	posts, err := pu.PostRepo.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pu *PostUsecase) GetPost(ctx context.Context, id int) (*models.Post, error) {
	return nil, nil
}

