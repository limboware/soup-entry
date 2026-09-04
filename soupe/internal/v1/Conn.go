package soupev1

import (
	limbov1 "github.com/limboware/limbo"
	queuev1 "github.com/limboware/pkg/structs/queue/v1"
)

var ConnType = limbov1.Compotype(0)

type Conn struct {
	Id limbov1.NetConn

	SessionID uint64

	ReadAES       []byte
	WriteAES      []byte
	WriteNonce    []byte
	ReadNonce     []byte
	ServerPubKey  []byte
	ServerPrivKey []byte
	PubKey        []byte
	Sharedkey     []byte

	RotateKey []byte
	KeyGen    uint8

	ReadBuffer        []byte
	ReadBufferPos     int
	ReadDesyncCounter uint8

	WriteBuffer    []byte
	WriteBufferId  string
	WriteBufferPos int

	IdleDuration int

	SeqInno  uint32
	SeqOutno uint32

	BytesIn  uint64
	BytesOut uint64

	OutcumQueue queuev1.Queue[*NetMessageEntry]

	CreatedAt    int64
	TouchedAt    int64
	HandshakedAt int64
	ConnectedAt  int64
}
