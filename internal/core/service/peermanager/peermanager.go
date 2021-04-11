package message

import (
	"context"

	"github.com/art-es/echpochmakcoin/internal/core/domain/message"
	"github.com/art-es/echpochmakcoin/internal/core/port"

	"google.golang.org/grpc/peer"
)

type PeerManager struct {
	serverStreams port.StreamRepository
	clientStreams port.StreamRepository
}

func New(serverStreams, clientStreams port.StreamRepository) *PeerManager {
	return &PeerManager{serverStreams, clientStreams}
}

func (m *PeerManager) AddServer(ctx context.Context, server port.Stream) error {
	return m.serverStreams.Add(ctx, server)
}

func (m *PeerManager) AddClient(ctx context.Context, client port.Stream) error {
	return m.clientStreams.Add(ctx, client)
}

func (m *PeerManager) SendAll(ctx context.Context, msg *message.Message) {
	from, ok := peer.FromContext(ctx)
	if !ok || from == nil {
		return
	}

	m.serverStreams.SendAll(from, msg)
	m.clientStreams.SendAll(from, msg)
}
