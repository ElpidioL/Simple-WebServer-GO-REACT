package PasswordHandler

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}
func CheckHash(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	if err != nil {
		return errors.New("Password Invalid.")
	}
	return err
}
