package soupev1

import (
	"time"

	"github.com/google/uuid"
	limbov1 "github.com/limboware/limbo"
	errnov1 "github.com/rejchev/errno"
)

type Ticket_t struct {
	Id          Ticket
	Token       string
	Requester   uint64
	ActivatedAt int64
	CreatedAt   int64
	Duration    int64
}

type TicketManager struct {
	tickets []Ticket

	gen  []uint8
	free []uint32

	token       []string
	requester   []uint64
	duration    []int64
	activatedAt []int64
	createdAt   []int64

	router map[Ticket]int
}

var tickets = TicketManager{}

func Tickets() *TicketManager {
	return &tickets
}

func (x *TicketManager) Init() errnov1.Code {
	x.tickets = make([]Ticket, 8)[:]

	x.gen = make([]uint8, 8)[:]
	x.free = make([]uint32, 8)[:]

	x.token = make([]string, 8)[:]
	x.requester = make([]uint64, 8)[:]

	x.duration = make([]int64, 8)[:]
	x.activatedAt = make([]int64, 8)[:]
	x.createdAt = make([]int64, 8)[:]

	x.router = make(map[Ticket]int)

	return errnov1.OK
}

func (x *TicketManager) Ticket(id Ticket, buff *Ticket_t) {
	*buff = Ticket_t{
		Id:          id,
		Token:       x.Token(id),
		Requester:   x.Requester(id),
		ActivatedAt: x.ActivatedAt(id),
		CreatedAt:   x.CreatedAt(id),
		Duration:    x.Duration(id),
	}
}

func (x *TicketManager) route(v Ticket) int {
	return x.router[v]
}

func (x *TicketManager) Activate(v Ticket) bool {
	if x.IsAlive(v) && x.activatedAt[v.Id()] == 0 {
		x.activatedAt[v.Id()] = time.Now().Unix()
		return true
	}

	return false
}

func (x *TicketManager) CreatedAt(v Ticket) int64 {
	return x.createdAt[x.route(v)]
}

func (x *TicketManager) Duration(v Ticket) int64 {
	return x.duration[x.route(v)]
}

func (x *TicketManager) ActivatedAt(v Ticket) int64 {
	return x.activatedAt[x.route(v)]
}

func (x *TicketManager) Requester(v Ticket) uint64 {
	return x.requester[x.route(v)]
}

func (x *TicketManager) Token(v Ticket) string {
	return x.token[x.route(v)]
}

func (x *TicketManager) IsAlive(v Ticket) bool {
	return x.gen[x.route(v)] == v.Gen()
}

func (x *TicketManager) Create(requester uint64, duration int64, buff *Ticket) bool {
	idx := uint32(0)

	if len(x.free) > 0 {
		idx = x.free[len(x.free)-1]
		x.free = x.free[:len(x.free)-1]
	} else {
		idx = uint32(len(x.gen))
		x.gen = append(x.gen, 0)
	}

	if len(x.token) <= int(idx) {
		x.token = append(x.token, "")
	}

	if len(x.requester) <= int(idx) {
		x.requester = append(x.requester, 0)
	}

	if len(x.activatedAt) <= int(idx) {
		x.activatedAt = append(x.activatedAt, 0)
	}

	if len(x.duration) <= int(idx) {
		x.duration = append(x.duration, 0)
	}

	if len(x.createdAt) <= int(idx) {
		x.createdAt = append(x.createdAt, 0)
	}

	x.token[idx] = uuid.NewString()
	x.requester[idx] = requester
	x.duration[idx] = duration
	x.createdAt[idx] = time.Now().Unix()

	*buff = MakeTicket(idx, x.gen[idx])

	x.tickets = append(x.tickets, *buff)
	x.router[*buff] = len(x.tickets) - 1

	limbov1.Events().Publish("tickets.created", *buff)

	return true
}

func (x *TicketManager) Iterator() *limbov1.Iterator[Ticket] {
	tickets := make([]Ticket, len(x.tickets))

	copy(tickets, x.tickets)

	return limbov1.NewIterator[Ticket](tickets)
}

// TODO: impl
func (x *TicketManager) Destroy(v Ticket) {}
