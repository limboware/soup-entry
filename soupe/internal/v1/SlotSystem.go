package soupev1

import (
	"time"
	"unsafe"

	limbov1 "github.com/limboware/limbo"
	errnov1 "github.com/rejchev/errno"
)

type SlotSystem struct{}

func slotSystem(buff *limbov1.System) {
	v := new(SlotSystem)

	*buff = limbov1.System{
		Instance:   unsafe.Pointer(v),
		Init:       v.Load,
		Activate:   nil,
		Update:     v.Update,
		Deactivate: nil,
		Destroy:    v.Unload,
	}
}

func (x *SlotSystem) Load() errnov1.Code {
	limbov1.Events().Subscribe("clients.connected", x.onClientConnected)
	limbov1.Events().Subscribe("clients.disconnect", x.onClientDisconnect)

	return errnov1.OK
}

func (x *SlotSystem) Unload() {}

func (x *SlotSystem) onClientConnected(_ string, data any) {
	if v, ok := data.(*ClientConnectedEvent); ok {
		Slots().Bind(uint8(v.Slot), v.SteamID)
	}
}

func (x *SlotSystem) onClientDisconnect(_ string, data any) {
	if v, ok := data.(*ClientDisconnectEvent); ok {
		Slots().Unbind(uint8(v.Slot))
	}
}

func (x *SlotSystem) Update(dt time.Duration) {
	x.rotate()
}

func (x *SlotSystem) rotate() {
	Slots().ForEach(func(v uint8) {
		ticket := Slots().Ticket(v)

		if Tickets().IsAlive(ticket) {
			if !Tickets().IsExpired(ticket) {
				return
			}

			Tickets().Remove(ticket)
		}

		if Tickets().CreateC(Slots().Owner(v), &ticket) {
			Slots().Rotate(v, ticket)
		}
	})
}
