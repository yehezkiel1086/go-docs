package repository

import (
	"context"
	"go-oauth/internal/adapter/storage/db/postgres"
	"go-oauth/internal/core/domain"
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

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *domain.User

	db := ur.db.GetDB()
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
