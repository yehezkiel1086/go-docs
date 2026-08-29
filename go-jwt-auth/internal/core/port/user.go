package port

import (
	"context"

	"github.com/yehezkiel1086/go-jwt-auth/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	UpdateUserByID(ctx context.Context, user *domain.User) error
	DeleteUserByID(ctx context.Context, id uint) error
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]domain.User, error)
	UpdateUserByID(ctx context.Context, id uint, user *domain.User) error
	DeleteUserByID(ctx context.Context, id uint) error
}