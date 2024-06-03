package users_repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/stretchr/testify/require"
)

func setupMocks(t *testing.T) (*users_repository.PostgreGetAccountByEmailRepository, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	getAccountByEmailRepository := &users_repository.PostgreGetAccountByEmailRepository{
		Db: db,
	}

	return getAccountByEmailRepository, mock, db
}

func TestGetAccountByEmail(t *testing.T) {
	repo, mock, db := setupMocks(t)
	defer db.Close()

	email := "test@example.com"
	query := regexp.QuoteMeta("SELECT * FROM users WHERE email = $1")

	rows := sqlmock.NewRows([]string{"id", "email", "password", "refresh_token"}).
		AddRow(1, email, "fake_hashed_password", "fake_refresh_token")

	mock.ExpectQuery(query).WithArgs(email).WillReturnRows(rows)

	account, err := repo.Get(email)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.Equal(t, email, account.Email)
}
