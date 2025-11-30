package repository

import (
	"context"
	"golang-redis/internal/entity"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UsersRepository struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewUsersRepository(db *gorm.DB, rdb *redis.Client) *UsersRepository {
	return &UsersRepository{DB: db, RDB: rdb}
}

func (u *UsersRepository) GetByEmail(ctx context.Context, email string) (*entity.MasterUsers, error) {
	var users entity.MasterUsers

	// todo: check from redis

	// check from database
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

	// todo: delete cache

	return users, nil
}
