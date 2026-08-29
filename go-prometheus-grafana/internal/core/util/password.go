package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
}

func ComparePassword(hashed, plain []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, plain)
}
