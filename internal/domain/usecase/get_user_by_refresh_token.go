package usecase

import "github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"

type GetUserByRefreshToken interface {
	Get(refreshToken string) (*models.AccountModel, error)
}
