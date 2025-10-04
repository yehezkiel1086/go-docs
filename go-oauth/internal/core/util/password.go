package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPwd), nil
}

func ComparePassword(hashedPwd string, pwd string) (error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)); err != nil {
		return err
	}

	return nil
}
