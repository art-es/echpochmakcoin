package port

import (
	"context"

	"github.com/art-es/echpochmakcoin/internal/core/domain/block"
)

type BlockchainStorage interface {
	ReadGenesisBlock(context.Context) (*block.Block, error)
	WriteBlock(context.Context, *block.Block) error
}
