package repository

import (
	"context"

	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/adapter/storage/postgres"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/domain"
)

type UserRepository struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.UserResponse, error) {
	db := ur.db.GetDB()
	if err := db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return &domain.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	db := ur.db.GetDB()

	var user domain.User
	if err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUsers(ctx context.Context) ([]domain.UserResponse, error) {
	db := ur.db.GetDB()

	var users []domain.UserResponse
	if err := db.WithContext(ctx).Model(&domain.User{}).Scan(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserByToken(ctx context.Context, token string) (*domain.User, error) {
	db := ur.db.GetDB()

	var user domain.User
	if err := db.WithContext(ctx).Where("confirmation_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := ur.db.GetDB()

	if err := db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
