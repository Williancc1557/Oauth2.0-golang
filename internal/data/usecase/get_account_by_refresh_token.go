package usecase

import (
	"github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
)

type DbGetAccountByRefreshToken struct {
	GetAccountByRefreshTokenRepository protocols.GetAccountByRefreshTokenRepository
}

func NewDbGetAccountByRefreshToken(GetAccountByRefreshTokenRepository protocols.GetAccountByRefreshTokenRepository) *DbGetAccountByRefreshToken {
	return &DbGetAccountByRefreshToken{
		GetAccountByRefreshTokenRepository,
	}
}

func (db DbGetAccountByRefreshToken) Get(refreshToken string) (*models.AccountModel, error) {
	account, err := db.GetAccountByRefreshTokenRepository.Get(refreshToken)

	if err != nil {
		return nil, err
	}

	return account, nil
}
