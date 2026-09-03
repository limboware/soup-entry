package soupev1

import (
	"crypto/rand"
	"math"
	"time"

	limbov1 "github.com/limboware/limbo"
	configsv2 "github.com/limboware/pkg/configs/v2"
	errnov1 "github.com/rejchev/errno"
)

type Ticket_t struct {
	Id          Ticket
	Token       []byte
	Requester   uint64
	ActivatedAt int64
	CreatedAt   int64
	Duration    int64
}

type TicketManager struct {
	token       [MAX_CLIENTS][16]byte
	requester   [MAX_CLIENTS]uint64
	createdAt   [MAX_CLIENTS]int64
	activatedAt [MAX_CLIENTS]int64
	gen         [MAX_CLIENTS]byte

	duration int64
	total    int8
}

var tickets = TicketManager{}

func Tickets() *TicketManager {
	return &tickets
}

func (x *TicketManager) Init() errnov1.Code {
	// TODO: add tickets_duration option to config
	if !configsv2.Int64Value("tickets_duration", &x.duration) {
		return errnov1.EINVAL
	}

	// TODO: add players option to config
	if !configsv2.IntKValue("players", 1, &x.total) {
		return errnov1.EINVAL
	}

	if x.total > MAX_CLIENTS {
		x.total = MAX_CLIENTS
	}

	x.token = [MAX_CLIENTS][16]byte{}
	x.requester = [MAX_CLIENTS]uint64{}
	x.createdAt = [MAX_CLIENTS]int64{}
	x.activatedAt = [MAX_CLIENTS]int64{}
	x.gen = [MAX_CLIENTS]byte{}

	return errnov1.OK
}

func (x *TicketManager) Ticket(id Ticket, buff *Ticket_t) {
	*buff = Ticket_t{
		Id:          id,
		Token:       x.Token(id),
		Requester:   x.Requester(id),
		ActivatedAt: x.ActivatedAt(id),
		CreatedAt:   x.CreatedAt(id),
		Duration:    x.Duration(),
	}
}

func (x *TicketManager) Activate(v Ticket) bool {
	if x.IsAlive(v) && x.activatedAt[v.Id()] == 0 {
		x.activatedAt[v.Id()] = time.Now().Unix()
		return true
	}

	return false
}

func (x *TicketManager) CreatedAt(v Ticket) int64 {
	return x.createdAt[v.Id()]
}

func (x *TicketManager) Duration() int64 {
	return x.duration
}

func (x *TicketManager) ActivatedAt(v Ticket) int64 {
	return x.activatedAt[v.Id()]
}

func (x *TicketManager) Requester(v Ticket) uint64 {
	return x.requester[v.Id()]
}

func (x *TicketManager) Token(v Ticket) []byte {
	return x.token[v.Id()][:]
}

func (x *TicketManager) Total() int8 {
	return x.total
}

func (x *TicketManager) IsAlive(v Ticket) bool {
	return x.gen[v.Id()] == v.Gen()
}

func (x *TicketManager) IsReserved(v Ticket) bool {
	return x.requester[v.Id()] != 0
}

// Reserved on duration time
func (x *TicketManager) Reserve(requester uint64, buff *Ticket) bool {
	if requester == 0 || buff == nil {
		return false
	}

	ticket := Ticket(0)

	if !x.First(&ticket, func(v Ticket) bool {
		return x.Requester(v) == 0
	}) {
		return false
	}

	x.requester[ticket.Id()] = requester

	*buff = ticket

	limbov1.Events().Publish("tickets.reserved", *buff)

	return true
}

func (x *TicketManager) IsExpired(v Ticket) bool {
	return x.activatedAt[v.Id()] == 0 && x.createdAt[v.Id()]+x.duration < time.Now().Unix()
}

func (x *TicketManager) First(buff *Ticket, cond func(Ticket) bool) bool {
	if buff == nil {
		return false
	}

	for i := range x.total {
		ticket := MakeTicket(uint8(i), x.gen[i])
		if cond(ticket) {
			*buff = ticket
			return true
		}
	}

	return false
}

func (x *TicketManager) Derive(v Ticket) {
	rand.Read(x.token[v.Id()][:])

	x.createdAt[v.Id()] = time.Now().Unix()
}

func (x *TicketManager) ForEach(fn func(Ticket)) {
	for i := range x.total {
		fn(MakeTicket(uint8(i), x.gen[i]))
	}
}

func (x *TicketManager) Unreserve(v Ticket) {
	if !x.IsAlive(v) {
		return
	}

	x.requester[v.Id()] = 0
	x.createdAt[v.Id()] = 0
}

func (x *TicketManager) Reset(v Ticket) {
	if !x.IsAlive(v) {
		return
	}

	limbov1.Events().Publish("tickets.reset", v)

	x.requester[v.Id()] = 0
	x.token[v.Id()] = [16]byte{}
	x.createdAt[v.Id()] = 0
	x.activatedAt[v.Id()] = 0

	if x.gen[v.Id()] == math.MaxUint8 {
		x.gen[v.Id()] = 0
	}

	x.gen[v.Id()]++

	limbov1.Events().Publish("tickets.reseted", v)
}
