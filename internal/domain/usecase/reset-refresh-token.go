package usecase

type ResetRefreshToken interface {
	Reset(userId string) (string, error)
}
