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
	Id        Ticket
	Token     []byte
	Requester limbov1.Entity
	Target    uint64
	CreatedAt int64
	Duration  int64
}

type TicketManager struct {
	container []Ticket

	token     [][16]byte
	requester []limbov1.Entity
	target    []uint64
	createdAt []int64
	duration  []int64

	gen  []byte
	free []uint16

	router map[Ticket]int

	baseDuration int64
}

var tickets = TicketManager{}

func Tickets() *TicketManager {
	return &tickets
}

func (x *TicketManager) Init() errnov1.Code {
	// TODO: add tickets_duration option to config
	if !configsv2.Int64Value("tickets.duration", &x.baseDuration) {
		return errnov1.EINVAL
	}

	x.container = make([]Ticket, 0, 8)
	x.token = make([][16]byte, 0, 8)
	x.requester = make([]limbov1.Entity, 0, 8)
	x.target = make([]uint64, 0, 8)
	x.createdAt = make([]int64, 0, 8)
	x.duration = make([]int64, 0, 8)
	x.gen = make([]byte, 0, 8)
	x.free = make([]uint16, 0, 8)

	x.router = map[Ticket]int{}

	return errnov1.OK
}

func (x *TicketManager) Ticket(v Ticket, buff *Ticket_t) {
	*buff = Ticket_t{
		Id:        v,
		Token:     x.Token(v),
		Target:    x.Target(v),
		Requester: x.Requester(v),
		Duration:  x.Duration(v),
		CreatedAt: x.CreatedAt(v),
	}
}

func (x *TicketManager) CreatedAt(v Ticket) int64 {
	return x.createdAt[v.Id()]
}

func (x *TicketManager) Duration(v Ticket) int64 {
	return x.duration[v.Id()]
}

func (x *TicketManager) Requester(v Ticket) limbov1.Entity {
	return x.requester[v.Id()]
}

func (x *TicketManager) Token(v Ticket) []byte {
	return x.token[v.Id()][:]
}

func (x *TicketManager) IsAlive(v Ticket) bool {
	return x.gen[v.Id()] == v.Gen()
}

func (x *TicketManager) Target(v Ticket) uint64 {
	return x.target[v.Id()]
}

// Create is creates new ticket
func (x *TicketManager) Create(requester limbov1.Entity, target uint64, duration int64, buff *Ticket) bool {
	if target == 0 || buff == nil {
		return false
	}

	if duration < 0 {
		duration = 0
	}

	idx := uint16(0)

	if len(x.free) > 0 {
		idx = x.free[len(x.free)-1]
		x.free = x.free[:len(x.free)-1]
	} else {
		if idx = uint16(len(x.gen)); idx == math.MaxUint16 {
			return false
		}

		x.gen = append(x.gen, 0)
	}

	if len(x.requester) <= int(idx) {
		x.requester = append(x.requester, 0)
	}

	if len(x.token) <= int(idx) {
		x.token = append(x.token, [16]byte{})
	}

	if len(x.target) <= int(idx) {
		x.target = append(x.target, 0)
	}

	if len(x.duration) <= int(idx) {
		x.duration = append(x.duration, 0)
	}

	if len(x.createdAt) <= int(idx) {
		x.createdAt = append(x.createdAt, 0)
	}

	if len(x.duration) <= int(idx) {
		x.duration = append(x.duration, 0)
	}

	rand.Read(x.token[idx][:])
	x.requester[idx] = requester
	x.target[idx] = target
	x.duration[idx] = duration
	x.createdAt[idx] = time.Now().Unix()

	*buff = MakeTicket(idx, x.gen[idx])

	innerIdx := len(x.container)

	x.container = append(x.container, *buff)

	x.router[*buff] = innerIdx

	limbov1.Events().Publish("tickets.new", *buff)

	return true
}

func (x *TicketManager) CreateA(requester limbov1.Entity, target uint64, buff *Ticket) bool {
	return x.Create(requester, target, x.baseDuration, buff)
}

func (x *TicketManager) CreateB(target uint64, duration int64, buff *Ticket) bool {
	return x.Create(SERVER, target, duration, buff)
}

func (x *TicketManager) CreateC(target uint64, buff *Ticket) bool {
	return x.CreateB(target, x.baseDuration, buff)
}

func (x *TicketManager) CreateD(requester limbov1.Entity, target uint64, buff *Ticket) bool {
	return x.Create(requester, target, x.baseDuration, buff)
}

func (x *TicketManager) IsExpired(v Ticket) bool {
	if x.Duration(v) == 0 {
		return false
	}

	return x.CreatedAt(v)+x.Duration(v) < time.Now().Unix()
}

func (x *TicketManager) route(v Ticket) int {
	if innerIdx, ok := x.router[v]; ok {
		return innerIdx
	}

	return -1
}

func (x *TicketManager) Remove(v Ticket) {
	if !x.IsAlive(v) {
		return
	}

	innerIdx := x.route(v)

	if innerIdx == -1 {
		return
	}

	limbov1.Events().Publish("tickets.remove", v)

	idx := v.Id()
	if (x.gen[idx] + 1) == math.MaxUint8 {
		x.gen[idx] = 0
	}

	x.gen[idx]++
	x.requester[idx] = 0
	x.target[idx] = 0
	x.token[idx] = [16]byte{}
	x.createdAt[idx] = 0
	x.duration[idx] = 0
	x.free = append(x.free, idx)

	len := len(x.container)

	if len > 1 && len-1 != innerIdx {
		x.container[innerIdx] = x.container[len-1]
		x.router[x.container[len-1]] = innerIdx
	}

	delete(x.router, v)
	x.container = x.container[:len-1]

	limbov1.Events().Publish("tickets.removed", v)
}

func (x *TicketManager) Iterator() *limbov1.Iterator[Ticket] {
	buff := make([]Ticket, len(x.container))

	copy(buff, x.container)

	return limbov1.NewIterator(buff)
}

func (x *TicketManager) Destroy() {}
