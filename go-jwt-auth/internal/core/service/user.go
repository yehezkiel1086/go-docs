package service

import (
	"context"

	"github.com/yehezkiel1086/go-jwt-auth/internal/core/domain"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/port"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	hashed, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) UpdateUserByID(ctx context.Context, id uint, input *domain.User) error {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		hashed, err := util.HashPassword(input.Password)
		if err != nil {
			return err
		}
		user.Password = hashed
	}

	return s.repo.UpdateUserByID(ctx, user)
}

func (s *UserService) DeleteUserByID(ctx context.Context, id uint) error {
	return s.repo.DeleteUserByID(ctx, id)
}
