package users_repository_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/stretchr/testify/require"
)

func setupMocks(t *testing.T) (*users_repository.GetAccountByEmailPostgreRepository, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	getAccountByEmailRepository := &users_repository.GetAccountByEmailPostgreRepository{
		Db: db,
	}

	return getAccountByEmailRepository, mock, db
}

func TestGetAccountByEmailPostgreRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
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
	})

	t.Run("UserNotFoundError", func(t *testing.T) {
		repo, mock, db := setupMocks(t)
		defer db.Close()

		email := "test@example.com"
		query := regexp.QuoteMeta("SELECT * FROM users WHERE email = $1")

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
		query := regexp.QuoteMeta("SELECT * FROM users WHERE email = $1")

		mock.ExpectQuery(query).WithArgs(email).WillReturnError(errors.New("fake-error"))

		account, err := repo.Get(email)
		require.Error(t, err)
		require.Nil(t, account)
	})
}
