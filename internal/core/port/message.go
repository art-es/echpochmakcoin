package port

import (
	"context"

	"github.com/art-es/echpochmakcoin/internal/core/domain/message"
	"google.golang.org/grpc/peer"
)

type MessageService interface {
	Handle(context.Context, *message.Message) error
}

type PeerManager interface {
	AddClient(ctx context.Context, client Stream) error
	AddServer(ctx context.Context, server Stream) error
	SendAll(ctx context.Context, msg *message.Message)
}

type Stream interface {
	Send(*message.Message) error
	Recv() (*message.Message, error)
}

type StreamRepository interface {
	Add(ctx context.Context, stream Stream) error
	Remove(addr string)
	SendAll(from *peer.Peer, msg *message.Message)
}
