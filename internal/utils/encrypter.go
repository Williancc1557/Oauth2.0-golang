package utils

import "golang.org/x/crypto/bcrypt"

type EncrypterUtil struct{}

func (e *EncrypterUtil) Hash(value string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (e *EncrypterUtil) CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
