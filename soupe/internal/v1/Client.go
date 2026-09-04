package soupev1

import (
	limbov1 "github.com/limboware/limbo"
)

var ClientType = limbov1.Compotype(0)

type Client struct {
	Id uint64

	SlotID uint8

	ConnectedAt int64
}
