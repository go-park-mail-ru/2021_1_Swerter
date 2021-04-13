package models

import (
	"context"
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	ID       int  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserID   int  `json:"userId"`
	FriendID int  `json:"friendId"`
	User     User `json:"user" gorm:"foreignKey:UserID"`
	Friend   User `json:"friend"  gorm:"foreignKey:FriendID"`
}

type FriendUsecase interface {
	AddFriend(c context.Context, session string, userFiend *User) error
	GetFriends(c context.Context, session string) ([]Friend, error)
}

type FriendRepository interface {
	GetFriends(ctx context.Context, userID int) ([]Friend, error)
	SaveFriend(ctx context.Context, userID int, fiendID int) error
}
