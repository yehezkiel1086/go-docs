package domain

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID uint `json:"user_id"`
	Role   Role `json:"role"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
