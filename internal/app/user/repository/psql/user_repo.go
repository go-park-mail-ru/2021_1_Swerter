package psql

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"mime/multipart"
	"my-motivation/internal/app/models"
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

func (urp *UserRepoPsql) SaveUser(c context.Context, u *models.User) error {
	timeout := time.Second * 2
	ctx, _ := context.WithTimeout(c, timeout)
	fmt.Printf("%+v\n", u)
	err := urp.DB.WithContext(ctx).Create(u).Error
	if err != nil {
		return err
	}
	return nil
}

func (urp *UserRepoPsql) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	//timeout := time.Second * 2
	//ctx, cancel := context.WithTimeout(c, time.Second * 2)
	u := models.User{}
	return &u, nil
}

func (urp *UserRepoPsql) GetUserById(ctx context.Context, id string) (*models.User, error) {
	//timeout := time.Second * 2
	//ctx, cancel := context.WithTimeout(c, time.Second * 2)
	u := models.User{}
	return &u, nil
}

func (urp *UserRepoPsql) GetPrivateUser(ctx context.Context, login string, password string) (*models.User, error) {
	//timeout := time.Second * 2
	//ctx, cancel := context.WithTimeout(c, time.Second * 2)
	u := models.User{}
	return &u, nil
}

func (urp *UserRepoPsql) UpdateUser(ctx context.Context, oldUser *models.User, newUser *models.User) error {
	//timeout := time.Second * 2
	//ctx, cancel := context.WithTimeout(c, time.Second * 2)
	return nil
}

func (urp *UserRepoPsql) UploadAvatar(c context.Context, user *models.User, file multipart.File) error {
	//timeout := time.Second * 2
	//ctx, cancel := context.WithTimeout(c, time.Second * 2)
	return nil
}
