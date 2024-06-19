package reset_refresh_token_repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type ResetRefreshTokenPostgreRepository struct {
	Db *sql.DB
}

func NewResetRefreshTokenPostgreRepository(db *sql.DB) *ResetRefreshTokenPostgreRepository {
	return &ResetRefreshTokenPostgreRepository{
		Db: db,
	}
}

func (rep *ResetRefreshTokenPostgreRepository) Reset(userId string) (string, error) {
	refreshToken := uuid.New().String()

	_, err := rep.Db.Exec("UPDATE users SET refresh_token = $1 WHERE id = $2", refreshToken, userId)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
