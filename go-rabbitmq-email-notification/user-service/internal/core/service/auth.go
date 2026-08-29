package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/domain"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/port"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/util"
	"gorm.io/gorm"
)

type AuthService struct {
	conf     *config.JWT
	userRepo port.UserRepository
}

func NewAuthService(conf *config.JWT, userRepo port.UserRepository) *AuthService {
	return &AuthService{
		conf,
		userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", domain.ErrUnauthorized
		}
		return "", fmt.Errorf("auth: failed to look up user: %w", err)
	}

	if !user.IsVerified {
		return "", domain.ErrUnauthorized
	}

	if err := util.CompareHashedPwd(user.Password, password); err != nil {
		return "", domain.ErrUnauthorized
	}

	token, err := util.GenerateJWTToken(as.conf, user)
	if err != nil {
		return "", fmt.Errorf("auth: failed to generate token: %w", err)
	}

	return token, nil
}
