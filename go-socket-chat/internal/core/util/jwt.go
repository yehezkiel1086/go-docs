package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-socket-chat/internal/adapter/config"
	"github.com/yehezkiel1086/go-socket-chat/internal/core/domain"
)

type JWTClaims struct {
	Email string `json:"email"`
	Username string `json:"username"`

	jwt.RegisteredClaims
}

func GenerateToken(conf *config.JWT, user *domain.User) (string, error) {
	mySigningKey := []byte(conf.Secret)

	claims := JWTClaims{
		Email: user.Email,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}
