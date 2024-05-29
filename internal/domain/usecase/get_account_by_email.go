package usecase

import "github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"

type GetAccountByEmail interface {
	Get(email string) (*models.AccountModel, error)
}
