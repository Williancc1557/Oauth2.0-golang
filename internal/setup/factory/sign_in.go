package factory

import (
	"database/sql"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/reset_refresh_token_repository"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/utils"
)

func MakeSignInController(db *sql.DB) *controllers.SignInController {
	encrypter := utils.NewEncrypterUtil()

	getAccountByEmailRepository := users_repository.NewGetAccountByEmailPostgreRepository(db)
	getAccountByEmail := usecase.NewGetAccountByEmail(getAccountByEmailRepository)
	resetRefreshToken := reset_refresh_token_repository.NewResetRefreshTokenPostgreRepository(db)

	return controllers.NewSignInController(getAccountByEmail, encrypter, resetRefreshToken)
}
