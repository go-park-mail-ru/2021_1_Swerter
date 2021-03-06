package usecase

import (
	"context"
	"mime/multipart"
	"my-motivation/internal/app/models"
	_sessionManager "my-motivation/internal/app/session/psql"
	"time"
)

type UserUsecase struct {
	userRepo       models.UserRepository
	postRepo       models.PostsRepository
	likeRepo       models.LikeRepository
	contextTimeout time.Duration
	sessionManager *_sessionManager.SessionsManagerPsql
}

func NewUserUsecase(u models.UserRepository, p models.PostsRepository, t time.Duration, sm *_sessionManager.SessionsManagerPsql, lr models.LikeRepository) models.UserUsecase {
	return &UserUsecase{
		userRepo:       u,
		postRepo:       p,
		likeRepo:       lr,
		contextTimeout: t,
		sessionManager: sm,
	}
}

func (uu *UserUsecase) GetFriends(c context.Context, session string) (map[string]*models.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userId, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return nil, err
	}

	user, err := uu.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	users, err := uu.userRepo.GetFriends(ctx, user)
	if err != nil {
		return nil, err
	}
	return users, nil
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

func (uu *UserUsecase) AddFriend(c context.Context, session string, userFiend *models.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	userId, err := uu.sessionManager.GetUserId(session)
	if err != nil {
		return err
	}

	user, err := uu.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	userFiend, err = uu.userRepo.GetUserById(ctx, userFiend.ID)
	if err != nil {
		return err
	}

	err = uu.userRepo.SaveFriend(ctx, user, userFiend)
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
