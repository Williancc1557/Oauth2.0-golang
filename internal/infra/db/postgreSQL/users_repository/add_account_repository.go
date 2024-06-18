package users_repository

import (
	"database/sql"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
	"github.com/google/uuid"
)

type AddAccountRepository struct {
	Db        *sql.DB
	Encrypter protocols.Encrypter
}

func (rep *AddAccountRepository) Add(data *protocols.AddAccountRepositoryInput) (*protocols.AddAccountRepositoryOutput, error) {
	query := "INSERT INTO users (id, email, password, refresh_token) VALUES ($1, $2, $3, $4)"
	userId := uuid.New().String()
	refreshToken := uuid.New().String()
	hashedPassword, err := rep.Encrypter.Hash(data.Password)

	if err != nil {
		return nil, err
	}

	_, err = rep.Db.Exec(query, userId, data.Email, hashedPassword, refreshToken)

	if err != nil {
		return nil, err
	}

	var account models.AccountModel
	err = rep.Db.QueryRow("SELECT * FROM users WHERE id = $1", userId).Scan(&account.Id, &account.Email, &account.Password, &account.RefreshToken)

	if err != nil {
		return nil, err
	}

	return &protocols.AddAccountRepositoryOutput{
		Id:           account.Id,
		Email:        account.Email,
		Password:     account.Password,
		RefreshToken: refreshToken,
	}, nil
}
