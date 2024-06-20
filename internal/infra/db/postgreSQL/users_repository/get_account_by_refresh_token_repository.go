package users_repository

import (
	"database/sql"
	"errors"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
)

type GetAccountByRefreshTokenPostgreRepository struct {
	Db *sql.DB
}

func NewGetAccountByRefreshTokenPostgreRepository(db *sql.DB) *GetAccountByRefreshTokenPostgreRepository {
	return &GetAccountByRefreshTokenPostgreRepository{
		Db: db,
	}
}

func (rep *GetAccountByRefreshTokenPostgreRepository) Get(refreshToken string) (*models.AccountModel, error) {
	var account models.AccountModel
	query := "SELECT * FROM users WHERE refresh_token = $1"
	err := rep.Db.QueryRow(query, refreshToken).Scan(&account.Id, &account.Email, &account.Password, &account.RefreshToken)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &account, nil
}
