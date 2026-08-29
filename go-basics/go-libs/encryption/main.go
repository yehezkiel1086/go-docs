package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func stdEncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func stdDecodeString(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

func urlEncode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

func urlDecode(encoded string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(encoded)
}

func sha1Encode(str string) (string) {
	sha := sha1.New()
	sha.Write([]byte(str))
	hash := sha.Sum(nil)
	hashedStr := fmt.Sprintf("%x", hash)

	return hashedStr
}

func saltedSha1Encode(str string) string {
	salt := time.Now().UnixNano()
	saltedStr := fmt.Sprintf("%s%d", str, salt)

	fmt.Println(saltedStr)

	sha := sha1.New()
	sha.Write([]byte(saltedStr))
	encoded := sha.Sum(nil)
	encodedStr := fmt.Sprintf("%x", encoded)

	return encodedStr
}

func encryptPassword(pass string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func comparePass(hash string, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)); err != nil {
		return err
	}
	return nil
}

func main() {
	encodedStr := stdEncodeString("Hello Base64!");
	fmt.Println(encodedStr)
	decodedStr, err := stdDecodeString(encodedStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decodedStr))

	url := "http://test.com"
	encodedUrl := urlEncode(url)
	fmt.Println(encodedUrl)

	decodedUrl, err := urlDecode(encodedUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(decodedUrl))

	encoded := sha1Encode("This is a secret")
	fmt.Println(encoded)

	encoded = saltedSha1Encode("This is a secret")
	fmt.Println(encoded)

	hashedPwd, err := encryptPassword("password")
	if err != nil {
		panic(err)
	}

	if err := comparePass(hashedPwd, "password"); err != nil {
		panic(err)
	} else {
		fmt.Println("Password matches")
	}
}
