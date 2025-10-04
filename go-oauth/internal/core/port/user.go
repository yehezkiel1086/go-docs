package port

import (
	"context"
	"go-oauth/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}

type UserService interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
}