package consensusvldr

import (
	"context"

	"github.com/art-es/echpochmakcoin/internal/core/domain/block"
)

type ConsensusValidator struct{}

func (c ConsensusValidator) ValidateBlock(context.Context, block.Block) error {
	return nil
}
