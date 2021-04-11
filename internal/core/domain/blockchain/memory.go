package blockchain

import (
	"sync"

	"github.com/art-es/echpochmakcoin/internal/core/domain/transaction"
)

// Memory represents a simple memory pool for unconfirmed transactions
type Memory struct {
	pool map[string]*transaction.Transaction
	sync.RWMutex
}

func NewMemory() *Memory {
	return &Memory{
		pool: make(map[string]*transaction.Transaction, 1024),
	}
}

func (m *Memory) Add(t *transaction.Transaction) {
	m.Lock()
	m.pool[t.HexID()] = t
	m.Unlock()
}
