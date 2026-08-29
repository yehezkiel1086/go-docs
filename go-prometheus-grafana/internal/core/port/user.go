package port

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
)

type UserRepository interface {
	CreateNewUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}

type UserService interface {
	CreateNewUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
	GetAllUsers(ctx context.Context) ([]domain.User, error)
}