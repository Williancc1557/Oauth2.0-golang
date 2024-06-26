package usecase

import (
	dataProtocols "github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/usecase"
)

type DbAddAccount struct {
	AddAccountRepository dataProtocols.AddAccountRepository
}

func NewDbAddAccount(AddAccountRepository dataProtocols.AddAccountRepository) *DbAddAccount {
	return &DbAddAccount{
		AddAccountRepository,
	}
}

func (db DbAddAccount) Add(account *usecase.AddAccountInput) (*usecase.AddAccountOutput, error) {
	accountData, err := db.AddAccountRepository.Add(&dataProtocols.AddAccountRepositoryInput{
		Email:    account.Email,
		Password: account.Password,
	})

	if err != nil {
		return nil, err
	}

	return &usecase.AddAccountOutput{
		Id:           accountData.Id,
		Email:        accountData.Email,
		Password:     accountData.Password,
		RefreshToken: accountData.RefreshToken,
	}, nil
}
