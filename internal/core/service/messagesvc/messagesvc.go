package messagesvc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/art-es/echpochmakcoin/internal/core/domain/block"
	"github.com/art-es/echpochmakcoin/internal/core/domain/message"
	"github.com/art-es/echpochmakcoin/internal/core/port"
)

type Service struct {
	PeerManager port.PeerManager
}

func (m Service) Handle(ctx context.Context, msg *message.Message) error {
	var (
		handlefunc func(context.Context, *message.Message) error
		errwrapmsg string
	)

	switch msg.Type {
	case message.AddBlock:
		handlefunc, errwrapmsg = m.handleAddBlock, "messagesvc.handleAddBlock"
	case message.AddTransaction:
		handlefunc, errwrapmsg = m.handleAddTransaction, "messagesvc.handleAddTransaction"
	default:
		return message.ErrUnknownType
	}

	if err := handlefunc(ctx, msg); err != nil {
		return fmt.Errorf("%s: %v", errwrapmsg, err)
	}
	m.PeerManager.SendAll(ctx, msg)
	return nil
}

func (m Service) handleAddBlock(ctx context.Context, msg *message.Message) error {
	var b block.Block
	if err := json.Unmarshal(msg.Data, &b); err != nil {
		return fmt.Errorf("json.Unmarshal: %v", err)
	}
	return nil
}

func (m Service) handleAddTransaction(ctx context.Context, msg *message.Message) error {
	return nil
}
