package models

import (
	"context"
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	ID         int  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserID     int  `json:"userId"`
	FriendID   int  `json:"friendId"`
	IsNotified bool `json:"isNotified"`
	User       User `json:"user" gorm:"foreignKey:UserID"`
	Friend     User `json:"friend"  gorm:"foreignKey:FriendID"`
}

type FriendUsecase interface {
	AddFriend(c context.Context, session string, userFiend *User) error
	GetFriends(c context.Context, session string) ([]User, error)
	GetFollowers(ctx context.Context, session string) ([]User, error)
	RemoveFriend(c context.Context, session string, removeFriend *User) error
}

type FriendRepository interface {
	GetFollowers(ctx context.Context, userID int) ([]Friend, error)
	GetSubscriptions(ctx context.Context, userID int) ([]Friend, error)
	SaveFriend(ctx context.Context, userID int, fiendID int) error
	RemoveFriend(ctx context.Context, userID int, friendID int) error
	GetFriend(ctx context.Context, userID int, friendID int) (*Friend, error)
	FriendRequestNotified(ctx context.Context, friend Friend) error
}
