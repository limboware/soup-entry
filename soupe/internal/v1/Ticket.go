package soupev1

import (
	"encoding/binary"
	"unsafe"
)

const TICKET_SIZE = 5

type Ticket [TICKET_SIZE]byte

func MakeTicket(id uint32, gen uint8) Ticket {
	ticket := Ticket{}

	ticket[4] = gen
	binary.LittleEndian.PutUint32(ticket[0:4], id)

	return ticket
}

func TicketFrom(v *byte) Ticket {
	tid := Ticket{}

	copy(tid[:], unsafe.Slice((*byte)(unsafe.Pointer(v)), TICKET_SIZE))

	return tid
}

func (x Ticket) Gen() uint8 {
	return x[4]
}

func (x Ticket) Id() uint32 {
	return binary.LittleEndian.Uint32(x[0:4])
}
