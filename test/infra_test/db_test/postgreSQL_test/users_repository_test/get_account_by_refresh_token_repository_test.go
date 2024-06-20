package users_repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/stretchr/testify/require"
)

func setupMocks(t *testing.T) (*users_repository.GetAccountByRefreshTokenPostgreRepository, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	getAccountByRefreshTokenRepository := users_repository.NewGetAccountByRefreshTokenPostgreRepository(db)

	return getAccountByRefreshTokenRepository, mock, db
}

func TestGetAccountByRefreshTokenPostgreRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repo, mock, db := setupMocks(t)
		defer db.Close()

		query := regexp.QuoteMeta("SELECT * FROM users WHERE refresh_token = $1")

		refreshToken := "fake-refresh-token"
		rows := sqlmock.NewRows([]string{"id", "email", "password", "refresh_token"}).
			AddRow(1, refreshToken, "fake_hashed_password", "fake_refresh_token")

		mock.ExpectQuery(query).WithArgs(refreshToken).WillReturnRows(rows)

		account, err := repo.Get(refreshToken)
		require.NoError(t, err)
		require.NotNil(t, account)
		require.Equal(t, refreshToken, account.Email)
	})

	t.Run("UserNotFoundError", func(t *testing.T) {
		repo, mock, db := setupMocks(t)
		defer db.Close()

		email := "test@example.com"
		query := regexp.QuoteMeta("SELECT * FROM users WHERE refresh_token = $1")

		nilRows := sqlmock.NewRows([]string{"id", "email", "password", "refresh_token"})

		mock.ExpectQuery(query).WithArgs(email).WillReturnRows(nilRows)

		account, err := repo.Get(email)
		require.Error(t, err)
		require.Nil(t, account)
	})

	t.Run("InvalidQueryError", func(t *testing.T) {
		repo, mock, db := setupMocks(t)
		defer db.Close()

		email := "test@example.com"
		query := regexp.QuoteMeta("SELECT * FROM users WHERE refresh_token = $1")

		mock.ExpectQuery(query).WithArgs(email).WillReturnError(errors.New("fake-error"))

		account, err := repo.Get(email)
		require.Error(t, err)
		require.Nil(t, account)
	})
}
