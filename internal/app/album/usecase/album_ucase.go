package usecase

import (
	"context"
	"mime/multipart"
	"my-motivation/internal/app/models"
	_sessionManager "my-motivation/internal/app/session/psql"
	"time"
)

type AlbumUsecase struct {
	UserRepo       models.UserRepository
	AlbumRepo       models.AlbumRepository
	contextTimeout time.Duration
	sessionManager *_sessionManager.SessionsManagerPsql
}

func NewAlbumUsecase(ur models.UserRepository, ar models.AlbumRepository, timeout time.Duration, sm *_sessionManager.SessionsManagerPsql) models.AlbumsUsecase {
	return &AlbumUsecase{
		UserRepo:       ur,
		AlbumRepo:       ar,
		contextTimeout: timeout,
		sessionManager: sm,
	}
}

func (au *AlbumUsecase) SaveAlbum(c context.Context, session string, fileHandlers map[string][]*multipart.FileHeader, album *models.Album) error {
	userId, err := au.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c, au.contextTimeout)
	defer cancel()

	userOwner, err := au.UserRepo.GetUserById(ctx, userId)
	if err != nil || userOwner == nil {
		return err
	}

	err = au.AlbumRepo.SaveAlbum(ctx, album, userOwner, fileHandlers)
	if err != nil {
		return err
	}
	return nil
}
