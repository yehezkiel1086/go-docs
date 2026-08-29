package port

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, id uint, req *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
}
