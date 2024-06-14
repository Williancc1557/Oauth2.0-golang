package usecase

import "github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"

type AddAccountInput struct {
	Email    string
	Password string
}

type AddAccount interface {
	Add(account AddAccountInput) (*models.AccountModel, error)
}
