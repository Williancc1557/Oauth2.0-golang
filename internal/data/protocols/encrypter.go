package protocols

type Encrypter interface {
	Hash(value string) (*string, error)
	Compare(value string, hashedValue string) bool
}
