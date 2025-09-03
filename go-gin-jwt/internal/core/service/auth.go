package service

import (
	"context"

	"github.com/yehezkiel1086/go-docs/go-gin-jwt/internal/core/domain"
	"github.com/yehezkiel1086/go-docs/go-gin-jwt/internal/core/port"
	"github.com/yehezkiel1086/go-docs/go-gin-jwt/internal/core/util"
)

type AuthService struct {
	repo port.UserRepository
}

func InitAuthService(repo port.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (svc *AuthService) Login(ctx context.Context, user *domain.User) (*domain.User, error) {
	// get request password
	reqPassword := user.Password

	// get by username
	user, err := svc.repo.GetUserByUsername(ctx, user.Username)
	if err != nil {
		return &domain.User{}, err
	}

	// compare password
	if err := util.ComparePassword(user.Password, reqPassword); err != nil {
		return &domain.User{}, err
	}

	return user, nil
}
