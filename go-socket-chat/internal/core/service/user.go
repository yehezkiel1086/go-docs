package service

import (
	"context"

	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/config"
	"github.com/yehezkiel1086/go-socket-chat/internal/core/domain"
	"github.com/yehezkiel1086/go-socket-chat/internal/core/port"
	"github.com/yehezkiel1086/go-socket-chat/internal/core/util"
)

type UserService struct {
	conf *config.JWT
	repo port.UserRepository
}

func NewUserService(conf *config.JWT, repo port.UserRepository) *UserService {
	return &UserService{
		conf,
		repo,
	}
}

func (s *UserService) Register(ctx context.Context, user *domain.User) (*domain.CreateUserRes, error) {
	// hash password
	hashed, err := util.HashPassword([]byte(user.Password))
	if err != nil {
		return nil, err
	}

	user.Password = string(hashed)

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) Login(ctx context.Context, creds *domain.LoginUserReq) (*domain.LoginUserRes, error) {
	// check email
	user, err := s.repo.GetUserByEmail(ctx, creds.Email)
	if err != nil {
		return nil, err
	}

	// check password
	if err := util.ComparePassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return nil, err
	}

	// generate access token
	token, err := util.GenerateToken(s.conf, user)
	if err != nil {
		return nil, err
	}

	return &domain.LoginUserRes{
		AccessToken: token,
		Username: user.Username,
	}, nil
}
