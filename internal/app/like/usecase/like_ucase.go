package usecase

import (
	"context"
	"my-motivation/internal/app/models"
	_sessionManager "my-motivation/internal/app/session/psql"
	"time"
)

type LikeUsecase struct {
	sessionManager *_sessionManager.SessionsManagerPsql
	likeRepo       models.LikeRepository
	contextTimeout time.Duration
}

func NewLikeUsecase (l models.LikeRepository, t time.Duration, sm *_sessionManager.SessionsManagerPsql) models.LikeUsecase {
	return &LikeUsecase{
		likeRepo:       l,
		contextTimeout: t,
		sessionManager: sm,
	}
}

func (lu *LikeUsecase) ChangeLike(ctx context.Context, session string, postID int) error {
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeout)
	defer cancel()

	userId, err := lu.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}
	isLiked, err := lu.likeRepo.IsLiked(ctx, userId, postID)
	if err != nil {
		return err
	}

	if isLiked == true {
		err = lu.likeRepo.DelLike(ctx, userId, postID)
	} else {
		err = lu.likeRepo.AddLike(ctx, userId, postID)
	}

	if err != nil {
		return err
	}

	return nil
}
