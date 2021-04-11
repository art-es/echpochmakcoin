package streamrepo

import (
	"context"
	"errors"
	"sync"

	"github.com/art-es/echpochmakcoin/internal/core/domain/message"
	"github.com/art-es/echpochmakcoin/internal/core/port"
	"google.golang.org/grpc/peer"
)

var ErrNoPeer = errors.New("peermanager: no peer")

type StreamRepository struct {
	M map[string]port.Stream
	sync.Mutex
}

func New() *StreamRepository {
	return &StreamRepository{M: make(map[string]port.Stream)}
}

func (s *StreamRepository) Add(ctx context.Context, stream port.Stream) error {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ErrNoPeer
	}

	s.Lock()
	s.M[p.Addr.String()] = stream
	s.Unlock()

	return nil
}

func (s *StreamRepository) Remove(addr string) {
	s.Lock()
	delete(s.M, addr)
	s.Unlock()
}

func (s *StreamRepository) SendAll(from *peer.Peer, msg *message.Message) {
	for addr, stream := range s.M {
		if addr == from.Addr.String() {
			continue
		}

		if err := stream.Send(msg); err != nil {
			s.Remove(addr)
		}
	}
}
