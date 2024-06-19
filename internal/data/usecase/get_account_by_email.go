package usecase

import (
	"github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
)

type DbGetAccountByEmail struct {
	GetAccountByEmailRepository protocols.GetAccountByEmailRepository
}

func NewGetAccountByEmail(GetAccountByEmailRepository protocols.GetAccountByEmailRepository) *DbGetAccountByEmail {
	return &DbGetAccountByEmail{
		GetAccountByEmailRepository,
	}
}

func (db *DbGetAccountByEmail) Get(email string) (*models.AccountModel, error) {
	value, err := db.GetAccountByEmailRepository.Get(email)

	if err != nil {
		return nil, err
	}

	return value, nil
}
