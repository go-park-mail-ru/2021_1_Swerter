package usecase

import (
	"context"
	"mime/multipart"
	"my-motivation/internal/app/models"
	_sessionManager "my-motivation/internal/app/session/psql"
	"time"
)

type IUserServiceController interface {
	Register(ctx context.Context, u *models.User) error
	Login(ctx context.Context, u *models.User) (string, error)
	GetUserBySession(ctx context.Context, session string) (*models.User, error)
	Logout(ctx context.Context, session string) error
}

type UserUsecase struct {
	//TODO:удалить репозиторий отсюда
	userRepo       models.UserRepository
	postRepo       models.PostsRepository
	albumRepo      models.AlbumRepository
	likeRepo       models.LikeRepository
	contextTimeout time.Duration
	userServiceApi IUserServiceController
	sessionManager *_sessionManager.SessionsManagerPsql
}

func NewUserUsecase(us IUserServiceController, u models.UserRepository, p models.PostsRepository, a models.AlbumRepository, t time.Duration, sm *_sessionManager.SessionsManagerPsql, lr models.LikeRepository) models.UserUsecase {
	return &UserUsecase{
		userServiceApi: us,
		userRepo:       u,
		postRepo:       p,
		albumRepo:      a,
		likeRepo:       lr,
		contextTimeout: t,
		sessionManager: sm,
	}
}

func (uu *UserUsecase) Register(c context.Context, u *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	err := uu.userServiceApi.Register(ctx, u)
	return err
}

func (uu *UserUsecase) Login(c context.Context, user *models.User) (string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	sess, err := uu.userServiceApi.Login(ctx, user)
	return sess, err
}

func (uu *UserUsecase) GetUserBySession(c context.Context, session string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	user, err := uu.userServiceApi.GetUserBySession(ctx, session)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu *UserUsecase) Logout(c context.Context, session string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	err := uu.userServiceApi.Logout(ctx, session)
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
	user.Albums, err = uu.albumRepo.GetUserAlbums(ctx, user)

	//Тут понять лайкнут ли и сколько всего лайков
	for i, _ := range user.Posts {
		isLiked, err := uu.likeRepo.IsLiked(ctx, user.ID, user.Posts[i].ID)
		if err != nil {
			return nil, err
		}
		user.Posts[i].Liked = isLiked
		likeCounter, err := uu.likeRepo.GetLikes(ctx, user.Posts[i].ID)
		if err != nil {
			return nil, err
		}
		user.Posts[i].LikeCounter = likeCounter
	}

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
