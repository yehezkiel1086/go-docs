package service

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/domain"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/port"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// hash password
	hashed, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashed

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, req *domain.User) (*domain.User, error) {
	// get user
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// update user
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		// hash password first
		hashed, err := util.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashed
	}

	return s.repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.DeleteUser(ctx, id)
}
