package port

type PrivateKey interface {
	PubKey() interface{}
	Sign([]byte) (Signature, error)
}

type Signature interface {
	Verify([]byte, interface{}) bool
}

type CryptographyService interface {
	GenerateKey() (PrivateKey, error)
	Sign(PrivateKey, []byte) (Signature, error)
}
