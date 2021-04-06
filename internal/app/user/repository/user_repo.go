package repository

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/hasher"
	"os"
)

type UserRepo struct {
	Users     map[string]*models.User
	IDToLogin map[string]string
	IDCounter int
}

func NewUserRepo() models.UserRepository {
	return &UserRepo{
		Users: map[string]*models.User{
			"AS@com.ru": {
				ID:        "id2",
				Login:     "AS@com.ru",
				Password:  "3339985ee8ca04466ce352663e4eadb95eb2e0a3f1565ba6d427d282266497dc",
				FirstName: "Artem",
				LastName:  "Sudorgin",
				Posts: make(map[int]*models.Post),
			},
		},
		IDCounter: 3,
		IDToLogin: map[string]string{
			"id2" : "AS@com.ru",
		},
	}
}

func (ur *UserRepo) SaveUser(ctx context.Context, u *models.User) error {
	if _, ok := ur.Users[u.Login]; ok {
		return errors.New(fmt.Sprintf("user %s was exist", u.Login))
	}

	u.ID = "id" + fmt.Sprint(ur.IDCounter)
	u.Password = hasher.Hash(u.Password)
	u.Posts = map[int]*models.Post{}

	ur.IDCounter++
	ur.Users[u.Login] = u
	ur.IDToLogin[u.ID] = u.Login

	return nil
}

func (ur *UserRepo) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {

	return &models.User{}, nil
}

func (ur *UserRepo) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user, ok := ur.Users[ur.IDToLogin[id]]
	if !ok {
		return nil, errors.New("no such user")
	}
	return user, nil
}


func (ur *UserRepo) GetPrivateUser(ctx context.Context, login string, password string) (*models.User, error) {
	if u, ok := ur.Users[login]; ok {
		if hasher.Hash(password) == u.Password {
			return u, nil
		}
	}
	return nil, errors.New("no user")
}

func (ur *UserRepo) UpdateUser(ctx context.Context, oldUser *models.User, newUser *models.User) error {
	if newUser.Password == "" {
		newUser.Password = oldUser.Password
	} else {
		if oldUser.Password != hasher.Hash(newUser.OldPassword) {
			return fmt.Errorf("not correct pass")
		}
		newUser.Password = hasher.Hash(newUser.Password)
	}

	newUser.ID = oldUser.ID

	if newUser.Login == "" {
		newUser.Login = oldUser.Login
	} else {
		ur.IDToLogin[newUser.ID] = newUser.Login
	}

	if newUser.FirstName == "" {
		newUser.FirstName = oldUser.FirstName
	}

	if newUser.LastName == "" {
		newUser.LastName = oldUser.LastName
	}

	newUser.Posts = oldUser.Posts
	newUser.Avatar = oldUser.Avatar

	delete(ur.Users, oldUser.Login)
	ur.Users[newUser.Login] = newUser
	return nil
}

func (ur *UserRepo) UploadAvatar(c context.Context, user *models.User, file multipart.File) error {

	user.Avatar = hasher.Hash(user.Login)
	ur.Users[user.Login] = user

	f, err := os.OpenFile("../../static/usersAvatar/"+user.Avatar, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, _ = io.Copy(f, file)
	return nil
}