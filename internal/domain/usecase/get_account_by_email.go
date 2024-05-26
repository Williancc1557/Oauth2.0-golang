package usecase

import "example/internal/domain/models"

type GetAccountByEmail interface {
	Get(email string) (*models.AccountModel, error)
}
