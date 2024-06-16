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
	AccessToken  string
	ExpiresIn    int
}

type AddAccountRepository interface {
	Add(account *AddAccountRepositoryInput) (*AddAccountRepositoryOutput, error)
}
