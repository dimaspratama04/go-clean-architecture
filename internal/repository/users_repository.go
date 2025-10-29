package repository

import (
	"context"
	"golang-redis/internal/entity"

	"gorm.io/gorm"
)

type UsersRepository struct {
	DB *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

func (u *UsersRepository) GetByEmail(ctx context.Context, email string) (*entity.MasterUsers, error) {
	var users entity.MasterUsers
	if err := u.DB.WithContext(ctx).Where("email = ?", email).First(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *UsersRepository) Create(ctx context.Context, username string, email string) (*entity.MasterUsers, error) {
	users := &entity.MasterUsers{
		Username: username,
		Email:    email,
	}

	if err := u.DB.WithContext(ctx).Create(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
