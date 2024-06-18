package users_repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Williancc1557/Oauth2.0-golang/internal/data/protocols"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/Williancc1557/Oauth2.0-golang/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	email          = "test@example.com"
	password       = "testpassword"
	hashedPassword = "fake_hashed_password"
)

func setupAddAccountRepositoryMocks(t *testing.T) (*users_repository.AddAccountRepository, sqlmock.Sqlmock, *sql.DB, *mocks.MockEncrypter, *gomock.Controller) {
	db, mock, err := sqlmock.New()
	ctrl := gomock.NewController(t)

	require.NoError(t, err)

	mockEncrypter := mocks.NewMockEncrypter(ctrl)

	addAccountRepository := &users_repository.AddAccountRepository{
		Db:        db,
		Encrypter: mockEncrypter,
	}

	return addAccountRepository, mock, db, mockEncrypter, ctrl
}

func TestAddAccountRepository(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		addAccountRepository, mock, db, mockEncrypter, ctrl := setupAddAccountRepositoryMocks(t)
		defer ctrl.Finish()
		defer db.Close()

		mockEncrypter.EXPECT().Hash(password).Return(hashedPassword, nil)

		queryExec := regexp.QuoteMeta("INSERT INTO users (id, email, password, refresh_token) VALUES ($1, $2, $3, $4)")
		mock.ExpectExec(queryExec).
			WithArgs(sqlmock.AnyArg(), email, hashedPassword, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		queryAddAccount := regexp.QuoteMeta("SELECT * FROM users WHERE id = $1")
		rows := sqlmock.NewRows([]string{"id", "email", "password", "refresh_token"}).
			AddRow("test", email, hashedPassword, "fake_refresh_token")
		mock.ExpectQuery(queryAddAccount).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

		input := &protocols.AddAccountRepositoryInput{
			Email:    email,
			Password: password,
		}
		account, err := addAccountRepository.Add(input)

		require.NoError(t, err)
		require.NotNil(t, account)
		require.Equal(t, email, account.Email)

	})
}
