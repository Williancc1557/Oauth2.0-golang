package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type CryptoUtil interface {
	GenerateFromPassword([]byte, int) ([]byte, error)
	CompareHashAndPassword(hashedPassword []byte, password []byte) error
}

type EncrypterUtil struct {
	Crypto CryptoUtil
}

type BcryptUtil struct{}

func (b *BcryptUtil) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, cost)
}

func (b *BcryptUtil) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func NewEncrypterUtil() *EncrypterUtil {
	return &EncrypterUtil{
		Crypto: &BcryptUtil{},
	}
}

func (e *EncrypterUtil) Hash(value string) (string, error) {
	hashed, err := e.Crypto.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (e *EncrypterUtil) Compare(value string, hashedValue string) bool {
	err := e.Crypto.CompareHashAndPassword([]byte(hashedValue), []byte(value))

	return err == nil
}
