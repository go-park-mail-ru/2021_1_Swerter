package psql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/hasher"
	"os"
	"time"
)

//SaveUser(ctx context.Context, u *User) error
//GetUserByLogin(ctx context.Context, login string) (*User, error)
//GetUserById(ctx context.Context, id string) (*User, error)
//GetPrivateUser(ctx context.Context, login string, password string) (*User, error)
//UpdateUser(ctx context.Context, oldUser *User, newUser *User) error
//UploadAvatar(c context.Context, user *User, file multipart.File) error

type UserRepoPsql struct {
	DB *gorm.DB
}

func NewUserRepoPsql(db *gorm.DB) *UserRepoPsql {
	return &UserRepoPsql{DB: db}
}

func (urp *UserRepoPsql) SaveUser(ctx context.Context, u *models.User) error {
	u.Password = hasher.Hash(u.Password)
	//u.Posts = []models.Post{{Text: "hello"},{Text: "world"}}
	err := urp.DB.WithContext(ctx).Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (urp *UserRepoPsql) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	u := models.User{}
	err := urp.DB.WithContext(ctx).First(&u, "login = ?", login).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (urp *UserRepoPsql) GetUserById(ctx context.Context, id int) (*models.User, error) {
	u := models.User{}
	err := urp.DB.WithContext(ctx).First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (urp *UserRepoPsql) GetPrivateUser(ctx context.Context, login string, password string) (*models.User, error) {
	u := models.User{}
	err := urp.DB.WithContext(ctx).First(&u, "login = ? AND password = ?", login, hasher.Hash(password)).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (urp *UserRepoPsql) UpdateUser(ctx context.Context, oldUser *models.User, newUser *models.User) error {
	if newUser.Password != "" {
		if oldUser.Password != hasher.Hash(newUser.OldPassword) {
			return errors.New("password not match")
		}
		newUser.Password = hasher.Hash(newUser.OldPassword)
	}

	err := urp.DB.WithContext(ctx).Model(oldUser).Updates(newUser).Error

	if err != nil {
		return err
	}
	return nil
}

func (urp *UserRepoPsql) UploadAvatar(ctx context.Context, u *models.User, file multipart.File) error {
	t := time.Now()
	salt := fmt.Sprintf(t.Format(time.RFC3339))

	u.Avatar = hasher.Hash(u.Login + salt)
	f, err := os.OpenFile("../../static/usersAvatar/"+u.Avatar, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	err = urp.DB.WithContext(ctx).Model(u).Update("avatar", u.Avatar).Error
	if err != nil {
		return err
	}

	return nil
}

func (urp *UserRepoPsql) GetFriends(ctx context.Context, user *models.User) (map[string]*models.User, error) {
	//if user == nil {
	//	return nil, errors.New("empty pointer to user")
	//}
	//
	//users := ur.Users[user.Login].Friends
	return nil, nil
}

func (urp *UserRepoPsql) SaveFriend(ctx context.Context, user *models.User, userFiend *models.User) error {
	//if user == nil || userFiend == nil {
	//	return errors.New("empty pointer to user")
	//}
	//
	//ur.Users[user.Login].Friends[userFiend.Login] = userFiend
	return nil
}
