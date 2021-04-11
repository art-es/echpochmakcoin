package grpc

import (
	"github.com/art-es/echpochmakcoin/internal/core/domain/message"
	"github.com/art-es/echpochmakcoin/internal/proto"
)

type ProtoStream interface {
	Send(*proto.Message) error
	Recv() (*proto.Message, error)
}

type Stream struct {
	Stream ProtoStream
}

func (s Stream) Send(msg *message.Message) error {
	return s.Stream.Send(&proto.Message{
		Type: msg.Type.String(),
		Data: msg.Data,
	})
}

func (s Stream) Recv() (*message.Message, error) {
	msg, err := s.Stream.Recv()
	if err != nil {
		return nil, err
	}
	return message.New(msg.Type, msg.Data)
}
