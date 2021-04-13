package usecase

import (
	"context"
	"errors"
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

func (uu *FriendUsecase) GetFriends(c context.Context, session string) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userID, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return nil, err
	}

	//На кого подписан
	userSubscribed, err := uu.friendRepo.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, err
	}

	//кто подписан на user
	userFollowers, err := uu.friendRepo.GetFollowers(ctx, userID)
	if err != nil {
		return nil, err
	}

	onlyFriends := []models.User{}
	for _, follower := range userFollowers {
		if uu.isFriend(follower.UserID, userSubscribed) {
			onlyFriends = append(onlyFriends, follower.User.Public())
		}
	}

	return onlyFriends, nil
}

func (uu *FriendUsecase) GetFollowers(c context.Context, session string) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userID, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return nil, err
	}

	//На кого подписан
	userSubscribed, err := uu.friendRepo.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, err
	}

	//кто подписан на тебя
	userFollowers, err := uu.friendRepo.GetFollowers(ctx, userID)
	if err != nil {
		return nil, err
	}

	onlyFollowers := []models.User{}
	for _, follower := range userFollowers {
		if !uu.isFriend(follower.UserID, userSubscribed) {
			follower.User.IsNotified = follower.IsNotified
			onlyFollowers = append(onlyFollowers, follower.User.Public())
		}
	}

	return onlyFollowers, nil
}

func (uu *FriendUsecase) AddFriend(c context.Context, session string, userFriend *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userID, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}

	//Проверяем добавляемого друга на существование
	userFriend, err = uu.userRepo.GetUserById(ctx, userFriend.ID)
	if err != nil || userFriend == nil {
		return err
	}

	if userFriend.ID == userID {
		return errors.New("can`t add to friend yourself")
	}

	err = uu.friendRepo.SaveFriend(ctx, userID, userFriend.ID)
	if err != nil {
		return err
	}

	return nil
}

func (uu *FriendUsecase) SearchFriend(c context.Context, session string, userPattern *models.User) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	if userPattern.FirstName == "" && userPattern.LastName == ""{
		return nil, errors.New("empty pattern. can`t find users")
	}

	if userPattern.FirstName != "" && userPattern.LastName == "" {
		users, err := uu.userRepo.SearchUsersByName(ctx, userPattern.FirstName)
		if err != nil  {
			return nil, err
		}

		if len(users) == 0 {
			users, err = uu.userRepo.SearchUsersBySurname(ctx, userPattern.FirstName)
		}
		if err != nil  {
			return nil, err
		}

		return users, nil
	}

	users, err := uu.userRepo.SearchUsersByFullName(ctx, userPattern.FirstName, userPattern.LastName)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uu *FriendUsecase) RemoveFriend(c context.Context, session string, removeFriend *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userID, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}

	friend, err := uu.friendRepo.GetFriend(ctx, removeFriend.ID, userID)
	if err != nil || friend == nil {
		return err
	}

	err = uu.friendRepo.FriendRequestNotified(ctx, *friend)
	if err != nil {
		return err
	}
	//delete this
	//friend, _ = uu.friendRepo.GetFriend(ctx, removeFriend.ID, userID)
	//fmt.Println(friend.IsNotified)

	//todo затестить
	err = uu.friendRepo.RemoveFriend(ctx, userID, removeFriend.ID)
	if err != nil {
		return err
	}

	return nil
}



func (uu *FriendUsecase) isFriend(friendID int, users []models.Friend) bool {
	for _, user := range users {
		if friendID == user.FriendID {
			return true
		}
	}
	return false
}
