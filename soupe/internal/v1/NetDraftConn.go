package soupev1

import limbov1 "github.com/limboware/limbo"

var NetDraftConnType = limbov1.Compotype(0)

type NetDraftConn struct {
	ReadBuffer    []byte
	ReadBufferPos int

	SleepTicks int

	Seqno uint32
	State int

	CreatedAt int64
}
