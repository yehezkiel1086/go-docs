package util

import (
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
