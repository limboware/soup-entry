package soupev1

import (
	"time"
	"unsafe"

	limbov1 "github.com/limboware/limbo"
	errnov1 "github.com/rejchev/errno"
)

type TicketSystem struct{}

func ticketSystem(buff *limbov1.System) {
	v := new(TicketSystem)

	*buff = limbov1.System{
		Instance:   unsafe.Pointer(v),
		Init:       v.Load,
		Activate:   nil,
		Update:     v.Update,
		Deactivate: nil,
		Destroy:    v.Unload,
	}
}

func (x *TicketSystem) Load() errnov1.Code {
	return errnov1.OK
}

func (x *TicketSystem) Unload() {}

func (x *TicketSystem) Update(dt time.Duration) {}
