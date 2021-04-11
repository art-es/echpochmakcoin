package btcd

import "github.com/btcsuite/btcd/btcec"

type Signature struct {
	*btcec.Signature
}

func (s Signature) Verify(hash []byte, pubkey interface{}) bool {
	btcecpubkey, ok := (pubkey).(*btcec.PublicKey)
	if !ok {
		return false
	}
	return s.Signature.Verify(hash, btcecpubkey)
}
