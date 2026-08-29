package service

import (
	"context"

	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/port"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/user-service/internal/core/util"
)

type AuthService struct {
	jwtConf  *config.JWT
	userRepo port.UserRepository
}

func NewAuthService(jwtConf *config.JWT, userRepo port.UserRepository) *AuthService {
	return &AuthService{
		jwtConf:  jwtConf,
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	// check email
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// check password
	if err := util.CheckPassword(password, user.Password); err != nil {
		return "", err
	}

	// generate jwt token
	token, err := util.GenerateJWT(s.jwtConf, user)
	if err != nil {
		return "", err
	}

	return token, nil
}
