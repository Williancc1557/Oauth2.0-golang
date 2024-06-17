package protocols

type AddAccountRepositoryInput struct {
	Email    string
	Password string
}

type AddAccountRepositoryOutput struct {
	Id           string
	Email        string
	Password     string
	RefreshToken string
}

type AddAccountRepository interface {
	Add(data *AddAccountRepositoryInput) (*AddAccountRepositoryOutput, error)
}
