package service

import (
	"context"
	"go-socket/internal/core/domain"
	"go-socket/internal/core/port"
	"go-socket/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func InitUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// hash password
	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return &domain.User{}, err
	}
	user.Password = hashedPwd

	// create user
	res, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return &domain.User{}, err
	}

	// return response
	return res, nil	
}
