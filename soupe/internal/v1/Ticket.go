package soupev1

import (
	"unsafe"
)

type Ticket uint16

func MakeTicket(id uint8, gen uint8) Ticket {
	return Ticket(uint16(gen)<<8 | uint16(id))
}

func TicketFrom(v *byte) Ticket {
	return MakeTicket(*v, *(*byte)(unsafe.Add(unsafe.Pointer(v), 1)))
}

func (x Ticket) Gen() uint8 {
	return uint8(x >> 8)
}

func (x Ticket) Id() uint8 {
	return uint8(x)
}
