package reset_refresh_token_repository_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/reset_refresh_token_repository"
	"github.com/stretchr/testify/require"
)

func setupMocks(t *testing.T) (*reset_refresh_token_repository.ResetRefreshTokenPostgreRepository, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	ResetRefreshTokenPostgreRepository := reset_refresh_token_repository.NewResetRefreshTokenPostgreRepository(db)

	return ResetRefreshTokenPostgreRepository, mock, db
}

func TestResetRefreshTokenPostgreRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repo, mock, db := setupMocks(t)
		defer db.Close()

		userId := "fake-user-id"
		query := regexp.QuoteMeta("UPDATE users SET refresh_token = $1 WHERE id = $2")

		mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), userId).WillReturnResult(sqlmock.NewResult(1, 1))

		refreshToken, err := repo.Reset(userId)
		require.NoError(t, err)
		require.NotEmpty(t, refreshToken)
	})

	t.Run("ErrorWhenExecuteQuery", func(t *testing.T) {
		repo, mock, db := setupMocks(t)
		defer db.Close()

		userId := "fake-user-id"
		query := regexp.QuoteMeta("UPDATE users SET refresh_token = $1 WHERE id = $2")

		mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), userId).WillReturnError(errors.New("fake-error"))

		refreshToken, err := repo.Reset(userId)
		require.Error(t, err)
		require.Empty(t, refreshToken)
	})
}
