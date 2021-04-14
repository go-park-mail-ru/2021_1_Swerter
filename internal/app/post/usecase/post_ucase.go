package usecase

import (
	"context"
	"mime/multipart"
	"my-motivation/internal/app/models"
	_sessionManager "my-motivation/internal/app/session/psql"
	"time"
)

type PostUsecase struct {
	UserRepo       models.UserRepository
	PostRepo       models.PostsRepository
	likeRepo       models.LikeRepository
	contextTimeout time.Duration
	sessionManager *_sessionManager.SessionsManagerPsql
}

func NewPostUsecase(ur models.UserRepository, pr models.PostsRepository, timeout time.Duration, sm *_sessionManager.SessionsManagerPsql, lr models.LikeRepository) models.PostsUsecase {
	return &PostUsecase{
		UserRepo:       ur,
		PostRepo:       pr,
		likeRepo:       lr,
		contextTimeout: timeout,
		sessionManager: sm,
	}
}

func (pu *PostUsecase) SavePost(c context.Context, session string, fileHandlers map[string][]*multipart.FileHeader, post *models.Post) error {
	userId, err := pu.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	userOwner, err := pu.UserRepo.GetUserById(ctx, userId)
	if err != nil || userOwner == nil {
		return err
	}

	err = pu.PostRepo.SavePost(ctx, post, userOwner, fileHandlers)
	if err != nil {
		return err
	}

	return nil
}

func (pu *PostUsecase) GetPosts(c context.Context, session string) ([]models.Post, error) {
	userId, err := pu.sessionManager.GetUserId(session)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	//TODO: доступ к страницам для авторизованных пользователей вынести в мидлвару
	_, err = pu.UserRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	posts, err := pu.PostRepo.GetPosts(ctx)
	for i, _ := range posts {
		isLiked, err := pu.likeRepo.IsLiked(ctx, userId, posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].Liked = isLiked
		likeCounter, err := pu.likeRepo.GetLikes(ctx, posts[i].ID)
		if err != nil {
			return nil, err
		}
		posts[i].LikeCounter = likeCounter
	}

	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pu *PostUsecase) GetPost(ctx context.Context, id int) (*models.Post, error) {
	return nil, nil
}
