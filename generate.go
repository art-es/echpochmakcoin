package echpochmakcoin

//go:generate protoc --go_out=plugins=grpc:. ./internal/proto/messagesvc.proto
//go:generate protoc --go_out=plugins=grpc:. ./internal/proto/paymentsvc.proto
