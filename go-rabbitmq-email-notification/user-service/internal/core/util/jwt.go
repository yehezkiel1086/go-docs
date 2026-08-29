package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/domain"
)

func GenerateJWTToken(conf *config.JWT, user *domain.User) (string, error) {
	mySigningKey := []byte(conf.Secret)

	duration, err := strconv.Atoi(conf.Duration)
	if err != nil {
		return "", fmt.Errorf("jwt: invalid duration config value %q: %w", conf.Duration, err)
	}

	now := time.Now()
	claims := domain.JWTClaims{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(duration) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("jwt: failed to sign token: %w", err)
	}

	return ss, nil
}
