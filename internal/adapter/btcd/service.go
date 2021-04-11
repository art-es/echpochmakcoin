package btcd

import (
	"github.com/art-es/echpochmakcoin/internal/core/port"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type Service struct{}

func (s Service) GenerateKey() (port.PrivateKey, error) {
	privkey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	privkey.PubKey().SerializeCompressed()
	return &PrivateKey{privkey}, nil
}

func (s Service) Sign(privkey port.PrivateKey, data []byte) (port.Signature, error) {
	hash := chainhash.DoubleHashB(data)
	return privkey.Sign(hash)
}
