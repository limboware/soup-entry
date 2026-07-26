package soupev1

import limbov1 "github.com/limboware/pkg/limbo/v1"

var TCPConnType = limbov1.Compotype(0)

type TCPConn struct {
	Id limbov1.ConnectionId

	Fd uintptr

	Frame            limbov1.NetMessage
	FramePos         int
	IsFrameIncoming  bool
	IsFrameFullyRead bool

	SleepTicks int

	ReadedAt int64
	WritedAt int64

	CreatedAt int64
}
