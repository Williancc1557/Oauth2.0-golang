package factory

import (
	"database/sql"

	"github.com/Williancc1557/Oauth2.0-golang/internal/data/usecase"
	"github.com/Williancc1557/Oauth2.0-golang/internal/infra/db/postgreSQL/users_repository"
	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/utils"
)

func MakeSignUpController(db *sql.DB) *controllers.SignUpController {
	createAccessToken := utils.NewCreateAccessTokenUtil()
	encrypterUtil := utils.NewEncrypterUtil()

	addAccountRepository := users_repository.NewAddAccountPostgreRepository(db, encrypterUtil)
	dbAddAccount := usecase.NewDbAddAccount(addAccountRepository)

	getAccountByEmailRepository := users_repository.NewGetAccountByEmailPostgreRepository(db)
	dbGetAccountByEmail := usecase.NewGetAccountByEmail(getAccountByEmailRepository)
	signUpController := controllers.NewSignUpController(dbGetAccountByEmail, dbAddAccount, createAccessToken)

	return signUpController
}
