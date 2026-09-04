package soupev1

import (
	"time"

	limbov1 "github.com/limboware/limbo"
	configsv2 "github.com/limboware/pkg/configs/v2"
	errnov1 "github.com/rejchev/errno"
)

type Slot_t struct {
	Id      uint8
	Ticket  Ticket
	Owner   uint64
	OwnedAt int64
}

type SlotManager struct {
	owner   [MAX_CLIENTS]uint64
	ticket  [MAX_CLIENTS]Ticket
	ownedAt [MAX_CLIENTS]int64

	total uint8
}

var slots = SlotManager{}

func Slots() *SlotManager {
	return &slots
}

func (x *SlotManager) Init() errnov1.Code {
	// TODO: add players option to config
	if !configsv2.IntKValue("players", 1, &x.total) {
		return errnov1.EINVAL
	}

	if x.total > MAX_CLIENTS {
		x.total = MAX_CLIENTS
	}

	x.owner = [MAX_CLIENTS]uint64{}
	x.ownedAt = [MAX_CLIENTS]int64{}
	x.ticket = [MAX_CLIENTS]Ticket{}

	return errnov1.OK
}

func (x *SlotManager) Slot(id uint8, buff *Slot_t) bool {
	if id >= x.total {
		return false
	}

	*buff = Slot_t{
		Id:      id,
		Ticket:  x.Ticket(id),
		Owner:   x.Owner(id),
		OwnedAt: x.OwnedAt(id),
	}

	return true
}

func (x *SlotManager) Duration(id uint8) int64 {
	if x.ownedAt[id] == 0 {
		return 0
	}

	return time.Now().Unix() - x.ownedAt[id]
}

func (x *SlotManager) Total() uint8 {
	return x.total
}

// Bind is bind slot with steamid64
func (x *SlotManager) Bind(id uint8, steamID uint64) bool {
	if id >= x.total || steamID == 0 || x.owner[id] != 0 {
		return false
	}

	x.owner[id] = steamID
	x.ownedAt[id] = time.Now().Unix()

	limbov1.Events().Publish("slots.binded", id)

	return true
}

func (x *SlotManager) Unbind(id uint8) {
	if id >= x.total || x.owner[id] == 0 {
		return
	}

	limbov1.Events().Publish("slots.unbind", id)

	x.owner[id] = 0
	x.ownedAt[id] = 0

	limbov1.Events().Publish("slots.unbinded", id)
}

func (x *SlotManager) Owner(id uint8) uint64 {
	return x.owner[id]
}

func (x *SlotManager) OwnedAt(id uint8) int64 {
	return x.ownedAt[id]
}

func (x *SlotManager) Ticket(id uint8) Ticket {
	return x.ticket[id]
}

func (x *SlotManager) IsAvailable(id uint8) bool {
	return x.Owner(id) == 0
}

func (x *SlotManager) Rotate(id uint8, ticket Ticket) {
	x.ticket[id] = ticket

	limbov1.Events().Publish("slots.rotated", id)
}

func (x *SlotManager) First(buff *uint8, cond func(uint8) bool) bool {
	if buff == nil {
		return false
	}

	for i := range x.total {
		if cond(i) {
			*buff = i
			return true
		}
	}

	return false
}

func (x *SlotManager) ForEach(fn func(uint8)) {
	for i := range x.total {
		fn(i)
	}
}

func (x *SlotManager) Reset(id uint8) {
	if id >= x.total {
		return
	}

	limbov1.Events().Publish("slots.reset", id)

	x.ownedAt[id] = 0
	x.owner[id] = 0
	x.ticket[id] = 0

	limbov1.Events().Publish("slots.reseted", id)
}
