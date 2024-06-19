package users_repository

import (
	"database/sql"
	"errors"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
)

type GetAccountByEmailPostgreRepository struct {
	Db *sql.DB
}

func NewGetAccountByEmailPostgreRepository(db *sql.DB) *GetAccountByEmailPostgreRepository {
	return &GetAccountByEmailPostgreRepository{
		Db: db,
	}
}

func (rep *GetAccountByEmailPostgreRepository) Get(email string) (*models.AccountModel, error) {
	var account models.AccountModel
	query := "SELECT * FROM users WHERE email = $1"
	err := rep.Db.QueryRow(query, email).Scan(&account.Id, &account.Email, &account.Password, &account.RefreshToken)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	if err != nil {
		return nil, err
	}

	return &account, nil
}
