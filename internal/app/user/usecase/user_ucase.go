package usecase

import (
	"context"
	"mime/multipart"
	"my-motivation/internal/app/models"
	"my-motivation/internal/app/session"
	"time"
)

type UserUsecase struct {
	userRepo       models.UserRepository
	postRepo       models.PostsRepository
	contextTimeout time.Duration
	sessionManager *session.SessionsManager
}

func NewUserUsecase(u models.UserRepository, p models.PostsRepository, t time.Duration, sm *session.SessionsManager) models.UserUsecase {
	return &UserUsecase{
		userRepo:       u,
		postRepo:       p,
		contextTimeout: t,
		sessionManager: sm,
	}
}

func (uu *UserUsecase) SaveUser(c context.Context, u *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err := uu.userRepo.SaveUser(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

//пока не используется на уровне usecase
func (uu *UserUsecase) GetPrivateUser(c context.Context, login string, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepo.GetPrivateUser(ctx, login, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu *UserUsecase) GetUserBySession(c context.Context, sessionValue string) (*models.User, error) {
	userId, err := uu.sessionManager.GetUserId(sessionValue)
	if err != nil {
		return nil, err
	}

	user, err := uu.GetUserById(c, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UserUsecase) LoginUser(c context.Context, user *models.User) (*models.Session, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	u, err := uu.userRepo.GetPrivateUser(ctx, user.Login, user.Password)
	if err != nil {
		return nil, err
	}

	sess, err := uu.sessionManager.Create(u.ID)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (uu *UserUsecase) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	return nil, nil
}

func (uu *UserUsecase) GetUserById(c context.Context, id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	user, err := uu.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	user.Posts, err = uu.postRepo.GetUserPosts(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *UserUsecase) UpdateUser(c context.Context, newUser *models.User, sessionId string) error {
	oldUser, err := uu.GetUserBySession(c, sessionId)
	if err != nil || oldUser == nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err = uu.userRepo.UpdateUser(ctx, oldUser, newUser)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUsecase) UploadAvatar(c context.Context, sessionId string, file multipart.File) error {
	user, err := uu.GetUserBySession(c, sessionId)

	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	err = uu.userRepo.UploadAvatar(ctx, user, file)

	if err != nil {
		return err
	}
	return nil
}

//Для логина и пароля
func (uu *UserUsecase) UpdateSecureUser(ctx context.Context, userId int, login string, pass string, newUser *models.User) error {
	return nil
}
