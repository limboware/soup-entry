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

func (x *TicketSystem) Update(dt time.Duration) {
	x.clean()

	x.update()
}

func (x *TicketSystem) clean() {
	Tickets().ForEach(func(v Ticket) {
		if Tickets().IsNotActivated(v) && Tickets().IsExpired(v) {
			Tickets().Reset(v)
		}
	})
}

func (x *TicketSystem) update() {
	Tickets().ForEach(func(v Ticket) {
		if Tickets().IsExpired(v) {
			Tickets().Rotate(v)
		}
	})
}
