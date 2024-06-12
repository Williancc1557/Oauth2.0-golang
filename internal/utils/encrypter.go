package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type EncrypterUtil struct{}

func (e *EncrypterUtil) Hash(value string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (e *EncrypterUtil) Compare(value string, hashedValue string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))

	return err == nil
}
