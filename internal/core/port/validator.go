package port

import (
	"context"

	"github.com/art-es/echpochmakcoin/internal/core/domain/block"
	"github.com/art-es/echpochmakcoin/internal/core/domain/transaction"
)

// ConsensusValidator is a set of consensus rules
type ConsensusValidator interface {
	// ValidateBlock checks the block
	// All requirements to block here:
	//	https://en.bitcoin.it/wiki/Protocol_rules#.22block.22_messages
	ValidateBlock(context.Context, block.Block) error

	// ValidateTransactions checks the transaction
	// All requirements to transaction here:
	//	https://en.bitcoin.it/wiki/Protocol_rules#.22tx.22_messages
	ValidateTransaction(context.Context, transaction.Transaction) error
}
