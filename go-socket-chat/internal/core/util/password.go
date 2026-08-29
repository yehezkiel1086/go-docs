package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
}

func ComparePassword(hashed, plain []byte) error {
	return bcrypt.CompareHashAndPassword(hashed, plain)
}
