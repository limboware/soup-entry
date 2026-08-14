package soupev1

import (
	limbov1 "github.com/limboware/limbo"
)

var ClientType = limbov1.Compotype(0)

type Client struct {
	Session limbov1.Entity

	Token string

	Version string

	ConnectedAt int64
}
