package soupev1

import limbov1 "github.com/limboware/limbo"

var NetConnType = limbov1.Compotype(0)

type NetConn struct {
	ConnId     limbov1.ConnectionId
	FileHandle uintptr
	
	CreatedAt  int64
}
