package users_repository

import (
	"database/sql"
	"errors"

	"github.com/Williancc1557/Oauth2.0-golang/internal/domain/models"
)

type PostgreGetAccountByEmailRepository struct {
	db *sql.DB
}

func (rep *PostgreGetAccountByEmailRepository) Get(email string) (*models.AccountModel, error) {
	var account models.AccountModel
	query := "SELECT * FROM users WHERE email = $1"
	err := rep.db.QueryRow(query, email).Scan(&account)

	if err == sql.ErrNoRows {
		return nil, errors.New("no account found")
	}

	if err != nil {
		return nil, err
	}

	return &account, nil
}
