package protocols

import "github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"

type GetAccountByRefreshTokenRepository interface {
	Get(refreshToken string) (*models.AccountModel, error)
}
