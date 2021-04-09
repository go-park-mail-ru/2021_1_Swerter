package models

import (
	"context"
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	ID     int  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	UserID int  `json:"userId"`
	PostID int  `json:"postId"`
	User   User `json:"user" gorm:"foreignKey:UserID"`
	Post   Post `json:"post"  gorm:"foreignKey:PostID"`
}

type LikeUsecase interface {
	ChangeLike(ctx context.Context, session string, postID int) error
}

type LikeRepository interface {
	AddLike(ctx context.Context, userID int, postID int) error
	DelLike(ctx context.Context, userID int, postID int) error
	IsLiked(ctx context.Context, userID int, postID int) (bool, error)
	GetLikes(ctx context.Context, userID int, postID int) (int, error)
}
