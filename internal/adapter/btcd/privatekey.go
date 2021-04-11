package btcd

import (
	"github.com/art-es/echpochmakcoin/internal/core/port"
	"github.com/btcsuite/btcd/btcec"
)

type PrivateKey struct {
	*btcec.PrivateKey
}

func (p PrivateKey) PubKey() interface{} {
	return p.PrivateKey.PubKey()
}

func (p PrivateKey) Sign(hash []byte) (port.Signature, error) {
	signature, err := p.PrivateKey.Sign(hash)
	if err != nil {
		return nil, err
	}
	return &Signature{signature}, nil
}
