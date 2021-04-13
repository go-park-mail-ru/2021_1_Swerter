package usecase

import (
	"context"
	"my-motivation/internal/app/models"
	_sessionManager "my-motivation/internal/app/session/psql"
	"time"
)

type FriendUsecase struct {
	friendRepo     models.FriendRepository
	userRepo       models.UserRepository
	contextTimeout time.Duration
	sessionManager *_sessionManager.SessionsManagerPsql
}

func NewFriendUsecase(f models.FriendRepository, u models.UserRepository, t time.Duration, sm *_sessionManager.SessionsManagerPsql) models.FriendUsecase {
	return &FriendUsecase{
		friendRepo:     f,
		userRepo:       u,
		contextTimeout: t,
		sessionManager: sm,
	}
}

func (uu *FriendUsecase) GetFriends(c context.Context, session string) ([]models.Friend, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userID, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return nil, err
	}

	friends, err := uu.friendRepo.GetFriends(ctx, userID)
	if err != nil {
		return nil, err
	}

	return friends, nil
}

func (uu *FriendUsecase) AddFriend(c context.Context, session string, userFiend *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userID, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}

	//Проверяем добавляемого друга на существование
	userFiend, err = uu.userRepo.GetUserById(ctx, userFiend.ID)
	if err != nil || userFiend == nil {
		return err
	}

	err = uu.friendRepo.SaveFriend(ctx, userID, userFiend.ID)
	if err != nil {
		return err
	}
	return nil
}
