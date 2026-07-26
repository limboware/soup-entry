package soupev1

import limbov1 "github.com/limboware/limbo"

var NetClientType = limbov1.Compotype(0)

type NetClient struct {
	ConnEnt limbov1.Entity

	ConnAddr string

	ReadBuffer    []byte
	ReadBufferPos int

	SleepTicks int

	CreatedAt int64
}
