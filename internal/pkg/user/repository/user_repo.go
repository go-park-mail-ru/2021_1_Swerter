package repository

import (
	"context"
	"my-motivation/internal/pkg/models"
)


type UserRepo struct {
	data map[string]*models.User
	IDCounter int
}

func NewUserRepo() models.UserRepository {
	return &UserRepo{
		data: map[string]*models.User{
			"AS@mail.ru": {
				ID:       "-1",
				Login:    "AS@mail.ru",
				Password: "AS890098",
				FirstName: "Artem",
				LastName: "Sudorgin",
			},
		},
	}
}

func (ur * UserRepo) SaveUser(ctx context.Context, u *models.User) error {
	return nil
}

func (ur * UserRepo) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	return &models.User{}, nil
}

func (ur * UserRepo) UpdateUser(ctx context.Context, userId int, newUser *models.User) error {
	return nil
}

