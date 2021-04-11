package message

import "errors"

// Message represents the object which transfer between nodes
type Message struct {
	Type Type
	Data []byte
}

func New(typ string, data []byte) (*Message, error) {
	t, err := parseMessageType(typ)
	if err != nil {
		return nil, err
	}
	return &Message{Type: t, Data: data}, nil
}

// Type represents the message type which defined the handle method of this message
type Type int

const (
	AddBlock Type = iota
	AddTransaction
)

var ErrUnknownType = errors.New("message: unknown type")

func (t Type) String() string {
	switch t {
	case AddBlock:
		return "addBlock"
	case AddTransaction:
		return "addTransaction"
	default:
		return ""
	}
}

func parseMessageType(s string) (Type, error) {
	switch s {
	case "addBlock":
		return AddBlock, nil
	case "addTransaction":
		return AddTransaction, nil
	default:
		return 0, ErrUnknownType
	}
}
