package factory

import (
	"database/sql"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/reset_refresh_token_repository"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/utils"
	"github.com/go-playground/validator/v10"
)

func MakeSignInController(db *sql.DB) *controllers.SignInController {
	encrypter := &utils.EncrypterUtil{}

	getAccountByEmailRepository := &users_repository.GetAccountByEmailPostgreRepository{
		Db: db,
	}
	getAccountByEmail := &usecase.DbGetAccountByEmail{
		GetAccountByEmailRepository: getAccountByEmailRepository,
	}
	resetRefreshToken := &reset_refresh_token_repository.ResetRefreshTokenPostgreRepository{
		Db: db,
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	return &controllers.SignInController{
		Encrypter:         encrypter,
		GetAccountByEmail: getAccountByEmail,
		ResetRefreshToken: resetRefreshToken,
		Validate:          validate,
	}
}
