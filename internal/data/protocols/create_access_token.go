package protocols

type CreateAccessTokenOutput struct {
	AccessToken string
	ExpiresIn   int
}

type CreateAccessToken interface {
	Create(userId string) (*CreateAccessTokenOutput, error)
}
