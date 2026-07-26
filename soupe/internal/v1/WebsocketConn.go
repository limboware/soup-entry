package soupev1

import limbov1 "github.com/limboware/pkg/limbo/v1"

var WebsocketConnType = limbov1.Compotype(0)

type WebsocketConn struct {
	Id limbov1.ConnectionId

	Message            limbov1.NetMessage
	MessageReadPos     int
	IsMessageIncoming  bool
	IsMessageFullyRead bool

	SleepTicks int

	ReadedAt int64
	WritedAt int64

	CreatedAt int64
}
