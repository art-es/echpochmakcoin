package btcd

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
)

type Keys struct {
	Private *btcec.PrivateKey
	Public  *btcec.PublicKey
	Address *btcutil.AddressPubKeyHash
}
