package protocols

type ResetRefreshTokenRepository interface {
	Reset(userId string) (string, error)
}
