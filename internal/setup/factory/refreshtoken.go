package factory

import (
	"database/sql"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/utils"
)

func MakeRefreshTokenController(db *sql.DB) *controllers.RefreshTokenController {
	createAccessToken := utils.NewCreateAccessTokenUtil()

	getAccountByRefreshTokenRepository := users_repository.NewGetAccountByRefreshTokenPostgreRepository(db)
	dbGetAccountByRefreshToken := usecase.NewDbGetAccountByRefreshToken(getAccountByRefreshTokenRepository)

	refreshTokenController := controllers.NewRefreshTokenController(dbGetAccountByRefreshToken, createAccessToken)

	return refreshTokenController
}
