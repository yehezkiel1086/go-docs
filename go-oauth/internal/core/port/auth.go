package port

import (
	"context"
	"go-oauth/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, req *domain.LoginReq) (*domain.User, error)
}