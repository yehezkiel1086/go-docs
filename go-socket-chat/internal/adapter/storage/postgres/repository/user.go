package repository

import (
	"context"

	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-socket-chat/internal/core/domain"
)

type UserRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.CreateUserRes, error) {
	db := r.db.GetDB()

	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return &domain.CreateUserRes{
		Email: user.Email,
		Username: user.Username,
	}, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	db := r.db.GetDB()

	var user domain.User

	if err := db.WithContext(ctx).First(user).Where("email", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
