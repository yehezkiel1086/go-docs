package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-jwt-auth/internal/adapter/config"
	"github.com/yehezkiel1086/go-jwt-auth/internal/core/domain"
)

func tokenDuration(tokenType string, conf *config.JWT) ([]byte, time.Duration, error) {
	switch tokenType {
	case "access":
		d, _ := strconv.Atoi(conf.AccessTokenDuration)
		return []byte(conf.AccessTokenSecret), time.Duration(d) * time.Minute, nil
	case "refresh":
		d, _ := strconv.Atoi(conf.RefreshTokenDuration)
		return []byte(conf.RefreshTokenSecret), time.Duration(d) * time.Hour * 24, nil
	default:
		return nil, 0, fmt.Errorf("invalid token type: %s", tokenType)
	}
}

func GenerateToken(tokenType string, user *domain.User, conf *config.JWT) (string, error) {
	token, _, err := GenerateTokenWithTTL(tokenType, user, conf)
	return token, err
}

func GenerateTokenWithTTL(tokenType string, user *domain.User, conf *config.JWT) (string, time.Duration, error) {
	signingKey, duration, err := tokenDuration(tokenType, conf)
	if err != nil {
		return "", 0, err
	}

	claims := &domain.JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(signingKey)
	return signed, duration, err
}

func ParseToken(tokenString string, secret string) (*domain.JWTClaims, error) {
	claims := &domain.JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}