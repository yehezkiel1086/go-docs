package service

import (
	"context"
	"go-oauth/internal/core/domain"
	"go-oauth/internal/core/port"
	"go-oauth/internal/core/util"
)

type AuthService struct {
	userRepo port.UserRepository
}

func InitAuthService(userRepo port.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, req *domain.LoginReq) (*domain.User, error) {
	// get user (check email)
	user, err := as.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &domain.User{}, err
	}

	// compare password (check password)
	if err := util.ComparePassword(user.Password, req.Password); err != nil {
		return &domain.User{}, err
	}

	// return logged in user
	return user, nil
}
