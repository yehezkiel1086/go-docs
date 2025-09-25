package port

import (
	"context"
	"go-oauth2/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)	
}
