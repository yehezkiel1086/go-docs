package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-grpc-inventory-microservices/services/product-service/internal/core/domain"
)

type JWTClaims struct {
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`

	jwt.RegisteredClaims
}

func GenerateJWT(conf *config.JWT, user *domain.User) (string, error) {
	mySigningKey := []byte(conf.Secret)

	duration, err := strconv.Atoi(conf.Duration)
	if err != nil {
		return "", err
	}

	// Create claims with multiple fields populated
	claims := &JWTClaims{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func ParseJWTToken(jwtSecret, tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("unknown claims type, cannot proceed")
}
