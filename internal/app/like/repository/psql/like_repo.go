package psql

import (
	"context"
	"gorm.io/gorm"
	"my-motivation/internal/app/models"
	"my-motivation/internal/pkg/utils/hasher"
)

type LikeRepoPsql struct {
	DB *gorm.DB
}

func NewLikeRepoPsql(db *gorm.DB) *LikeRepoPsql {
	return &LikeRepoPsql{DB: db}
}
func (lr *LikeRepoPsql) AddLike(ctx context.Context, userID int, postID int) error {
	like := models.Like{
		UserID: userID,
		PostID: postID,
	}

	err := lr.DB.WithContext(ctx).Create(&like).Error

	if err != nil {
		return err
	}

	return nil
}

func (lr *LikeRepoPsql) DelLike(ctx context.Context, userID int, postID int) error {
	//TODO::del like
	return nil
}

func (lr *LikeRepoPsql) IsLiked(ctx context.Context, userID int, postID int) (bool, error) {
	l := models.Like{}
	err := lr.DB.WithContext(ctx).First(&l, "user_id = ? AND post_id = ?", userID, postID).Error
	if err != nil {
		return nil, err
	}
	//TODO:Проверка на найденность
	if {
		return false, err
	}
	return true, nil
}

func (lr *LikeRepoPsql) GetLikes(ctx context.Context, postID int) (int, error) {
	var likes []models.Like
	err := lr.DB.WithContext(ctx).Where("post_id = ?", postID).Find(&likes).Error
	if err != nil {
		return 0,err
	}
	return len(likes), nil
}
