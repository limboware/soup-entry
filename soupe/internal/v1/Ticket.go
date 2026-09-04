package soupev1

type Ticket uint32

func MakeTicket(id uint16, gen uint8) Ticket {
	return Ticket(uint32(gen)<<16 | uint32(id))
}

func (x Ticket) Gen() uint8 {
	return uint8(x >> 16)
}

func (x Ticket) Id() uint16 {
	return uint16(x)
}
