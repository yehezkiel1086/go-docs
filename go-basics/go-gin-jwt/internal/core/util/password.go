package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}

func ComparePassword(hashedPwd string, password string) (error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password)); err != nil {
		return err
	}

	return nil
}
