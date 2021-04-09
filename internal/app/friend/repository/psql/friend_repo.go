package psql

import (
	"context"
	"gorm.io/gorm"
	"my-motivation/internal/app/models"
)

type FriendRepoPsql struct {
	DB *gorm.DB
}

func NewFriendRepoPsql(db *gorm.DB) *FriendRepoPsql {
	return &FriendRepoPsql{DB: db}
}

func (frp *FriendRepoPsql) GetSubscriptions(ctx context.Context, userID int) ([]models.Friend, error) {
	user := []models.Friend{}
	err := frp.DB.WithContext(ctx).Preload("User").Preload("Friend").Find(&user, "user_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (frp *FriendRepoPsql) GetFollowers(ctx context.Context, userID int) ([]models.Friend, error) {
	user := []models.Friend{}
	err := frp.DB.WithContext(ctx).Preload("User").Preload("Friend").Find(&user, "friend_id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (frp *FriendRepoPsql) SaveFriend(ctx context.Context, userID int, friendID int) error {
	friend := models.Friend{}
	friend.UserID = userID
	friend.FriendID = friendID
	err := frp.DB.WithContext(ctx).Create(&friend).Error
	if err != nil {
		return err
	}

	return nil
}
