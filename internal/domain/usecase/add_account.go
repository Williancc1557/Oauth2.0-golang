package usecase

type AddAccountInput struct {
	Email    string
	Password string
}

type AddAccountOutput struct {
	Id           string
	Email        string
	Password     string
	RefreshToken string
}

type AddAccount interface {
	Add(account *AddAccountInput) (*AddAccountOutput, error)
}
