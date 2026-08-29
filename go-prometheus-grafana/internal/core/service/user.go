package service

import (
	"context"

	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/domain"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/port"
	"github.com/yehezkiel1086/go-prometheus-grafana/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

func (s *UserService) CreateNewUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashed, err := util.HashPassword([]byte(user.Password))
	if err != nil {
		return nil, err
	}

	user.Password = string(hashed)

	return s.repo.CreateNewUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return s.repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.GetAllUsers(ctx)
}
