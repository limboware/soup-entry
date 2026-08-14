package soupev1

import (
	limbov1 "github.com/limboware/limbo"
	queuev1 "github.com/limboware/pkg/structs/queue/v1"
)

var NetClientType = limbov1.Compotype(0)

type NetClient struct {
	Id uint64

	ConnId         limbov1.NetConn
	ConnFileHandle uintptr
	ConnAddr       uint64

	Slot limbov1.Entity

	SteamId uint64

	ConnReadAES       []byte
	ConnWriteAES      []byte
	ConnWriteNonce    []byte
	ConnReadNonce     []byte
	ConnChaChaKey     []byte
	ConnServerPubKey  []byte
	ConnServerPrivKey []byte
	ConnPubKey        []byte
	ConnSharedkey     []byte

	ConnRotateKey []byte
	ConnKeyGen    uint8

	ReadBuffer        []byte
	ReadBufferPos     int
	ReadDesyncCounter uint8
	// ReadSleepDuration int

	WriteBuffer    []byte
	WriteBufferId  string
	WriteBufferPos int
	// WriteSleepDuration int

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
