package service

import (
	"context"

	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/core/domain"
	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/core/port"
	"github.com/yehezkiel1086/go-docs/go-echo-rest/internal/core/util"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) RegisterNewUser(ctx context.Context, user domain.User) (domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	user.Password = hashedPassword
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *UserService) UpdateUser(ctx context.Context, id int64, user *domain.User) (domain.User, error) {
	// get user
	existingUser, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	// update user
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {
			return domain.User{}, err
		}
		existingUser.Password = hashedPassword
	}

	return s.repo.UpdateUser(ctx, existingUser)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int) ([]domain.User, error) {
	return s.repo.ListUsers(ctx, limit, offset)
}

func (s *UserService) CountUsers(ctx context.Context, search string) (int64, error) {
	return s.repo.CountUsers(ctx, search)
}
