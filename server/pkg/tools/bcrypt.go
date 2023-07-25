package tools

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(pwd string) (string, error) {
	if hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(hash), nil
	}
}

func ComparePassword(hash, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)); err != nil {
		return false
	}
	return true
}
