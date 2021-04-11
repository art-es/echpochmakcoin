package grpc

import (
	"fmt"
	"io"

	"github.com/art-es/echpochmakcoin/internal/core/domain/message"
	"github.com/art-es/echpochmakcoin/internal/core/port"
	"github.com/art-es/echpochmakcoin/internal/proto"

	"github.com/sirupsen/logrus"
)

type MessageService struct {
	PeerManager port.PeerManager
	Service     port.MessageService
}

func (a MessageService) Broadcast(srv proto.MessageService_BroadcastServer) error {
	if err := a.PeerManager.AddServer(srv.Context(), &Stream{srv}); err != nil {
		return err
	}

	for {
		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("MessageService_BroadcastServer.Recv: %v", err)
		}

		msg, err := message.New(in.Type, in.Data)
		if err != nil {
			return err
		}

		if err = a.Service.Handle(srv.Context(), msg); err != nil {
			logrus.WithError(err).Error("adapter: messagesvc: GRPCAdapter.Broadcast")
		}
	}
}
