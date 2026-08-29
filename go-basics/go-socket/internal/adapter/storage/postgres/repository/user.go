package repository

import (
	"context"
	storage "go-socket/internal/adapter/storage/postgres"
	"go-socket/internal/core/domain"
)

type UserRepository struct {
	db *storage.DB
}

func InitUserRepository(db *storage.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := ur.db.GetDB()
	if err := db.Create(&user).Error; err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
