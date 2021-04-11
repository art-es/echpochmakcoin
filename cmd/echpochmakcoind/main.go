package main

import (
	"net"
	"os"

	"github.com/art-es/echpochmakcoin/cmd"
	grpcadapter "github.com/art-es/echpochmakcoin/internal/adapter/grpc"
	"github.com/art-es/echpochmakcoin/internal/proto"

	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := cli.NewApp()
	app.Name = cmd.ACIILogo + "\nechpochmakcoind"
	app.Version = "0.0.1"
	app.Usage = "Run EchpochmakCoin as a daemon"
	app.Action = func(*cli.Context) error { return run() }
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.WithError(err).Fatal("main: net.Listen")
	}

	server := grpc.NewServer()
	proto.RegisterMessageServiceServer(server, grpcadapter.MessageService{})

	return server.Serve(lis)
}
