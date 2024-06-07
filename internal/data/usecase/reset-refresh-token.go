package usecase

import "github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"

type DbResetRefreshToken struct {
	ResetRefreshTokenRepository protocols.ResetRefreshTokenRepository
}

func (rep DbResetRefreshToken) Reset(userId string) (string, error) {
	return rep.ResetRefreshTokenRepository.Reset(userId)
}
