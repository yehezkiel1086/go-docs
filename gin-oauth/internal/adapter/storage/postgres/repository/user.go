package repository

import (
	"context"
	"go-oauth2/internal/adapter/storage/postgres"
	"go-oauth2/internal/core/domain"
)

type UserRepository struct {
	db *postgres.DB
}

func InitUserRepository(db *postgres.DB) *UserRepository {
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
