package block

import "encoding/hex"

type Block struct {
	Header
	Hash []byte
}

func (b Block) HexHash() string {
	return hex.EncodeToString(b.Hash)
}

func New() *Block {
	return &Block{}
}

type Header struct {
	PreviousBlockHash []byte
	MercleRootHash    []byte
	Timestamp         int64
	Target            string
	Nonce             string
}
