package soupev1

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"hash/crc32"
	"net"
	"slices"
	"syscall"
	"time"
	"unsafe"

	"github.com/google/uuid"
	limbov1 "github.com/limboware/limbo"
	configsv2 "github.com/limboware/pkg/configs/v2"
	loggerv1 "github.com/limboware/pkg/logger/v1"
	pbsoupev1 "github.com/limboware/pkg/proto/soupe/v1"
	queuev1 "github.com/limboware/pkg/structs/queue/v1"
	sysutils "github.com/limboware/soupe/utils/syscalls"
	errnov1 "github.com/rejchev/errno"
	"google.golang.org/protobuf/proto"
)

var TcpConnType = limbov1.NetConnType(0)

type NetMessageEntry struct {
	Id        string
	Typ       limbov1.NetMessageType
	Payload   []byte
	Any       []byte
	CreatedAt time.Time
}

type NetSystem struct {
	listener       *net.TCPListener
	maxConnections uint32
}

func netSystem(buff *limbov1.System) {
	v := new(NetSystem)

	*buff = limbov1.System{
		Instance:   unsafe.Pointer(v),
		Init:       v.Load,
		Activate:   v.Activate,
		Update:     v.Update,
		Deactivate: v.Deactivate,
		Destroy:    v.Unload,
	}
}

// Activate implements [limbov1.ISystem].
func (x *NetSystem) Activate() bool {
	return true
}

// Deactivate implements [limbov1.ISystem].
func (x *NetSystem) Deactivate() {
	if x.listener != nil {
		_ = x.listener.Close()
	}
}

// Load implements [limbov1.ISystem].
func (x *NetSystem) Load() errnov1.Code {
	if !limbov1.NetConnTypes().Register(&TcpConnType) {
		return errnov1.ECALL
	}

	if !configsv2.IntKValue("maxconns", 32, &x.maxConnections) {
		x.maxConnections = 512
	}

	listener, err := net.Listen("tcp", ":"+configsv2.Get().ValueA("port", "27720"))

	if err != nil {
		loggerv1.Get().System("server err on setup listener (err: %v)", err)
		return errnov1.EINVAL
	}

	if tcpListener, ok := listener.(*net.TCPListener); ok {
		x.listener = tcpListener
	}

	if x.listener == nil {
		listener.Close()
		return errnov1.EINVAL
	}

	if rawConn, err := x.listener.SyscallConn(); true {
		if err != nil {
			loggerv1.Get().System("server err on setup listener control (err: %v)", err)
			x.listener.Close()
			return errnov1.ECALL
		}

		rawConn.Control(func(fd uintptr) {
			sysutils.SetNonblock(fd, true)
		})
	}

	return errnov1.OK
}

// Unload implements [limbov1.ISystem].
func (x *NetSystem) Unload() {
	if x.listener != nil {
		x.listener.Close()
	}
}

// Update implements [limbov1.ISystem].
func (x *NetSystem) Update(dt time.Duration) {
	x.Push()

	x.Cleanup()

	x.Accept()
}

func (x *NetSystem) Cleanup() {
	limbov1.Components().IterateB(NetClientType, func(e limbov1.Entity, p unsafe.Pointer) {
		if limbov1.ContainsComponent(e, DestroyType) {
			return
		}

		sptr := (*NetClient)(p)

		if sptr == nil {
			NetDisconnect(e, "diconnect")
			return
		}

		tn := time.Now().Unix()

		if sptr.SteamId != 0 {
			return
		}

		// too long handshake
		if sptr.HandshakedAt == 0 {
			if tn-sptr.CreatedAt >= 15 {
				NetDisconnect(e, "timexpired")
			}

			return
		}
	})
}

func (x *NetSystem) Push() {
	limbov1.Components().IterateB(NetClientType, func(e limbov1.Entity, p unsafe.Pointer) {
		sptr := (*NetClient)(p)

		if !limbov1.Networks().IsAlive(sptr.ConnId) || sptr.ConnFileHandle == 0 {
			return
		}

		if sptr.WriteBufferId == "" {
			if sptr.OutcumQueue.Len() == 0 {
				return
			}

			netMsgEntry := sptr.OutcumQueue.Peak()

			plen := len(netMsgEntry.Payload)

			total := limbov1.NETMSG_HEADER_SIZE + plen

			if sptr.WriteBuffer == nil {
				sptr.WriteBuffer = make([]byte, total)
			}

			if cap(sptr.WriteBuffer) < total {
				sptr.WriteBuffer = slices.Grow(sptr.WriteBuffer, total-cap(sptr.WriteBuffer))
			}

			nmsg := (*limbov1.NetMessage)(&sptr.WriteBuffer[0])

			builder := nmsg.Builder().
				PutKeyID(sptr.ConnKeyGen).
				PutType(netMsgEntry.Typ.Closer()).
				PutTimestamp(uint64(netMsgEntry.CreatedAt.UnixMilli())).
				PutSequence(sptr.SeqOutno).
				PutSessionID(sptr.Id)

			if len(netMsgEntry.Any) != 0 {
				builder.ReserveBytes(0, &netMsgEntry.Any[0], len(netMsgEntry.Any))
			}

			if pbsoupev1.MsgType(netMsgEntry.Typ) != pbsoupev1.MsgType_Welcome {
				builder.PutFlags(limbov1.FlagEncrypted)
			}

			pl := netMsgEntry.Payload

			if plen != 0 {
				if nmsg.Flags().Contains(limbov1.FlagEncrypted) {
					nonce := sptr.ConnWriteNonce[0:]
					binary.LittleEndian.PutUint32(nonce[8:], binary.LittleEndian.Uint32(nonce[8:])^nmsg.Sequence())

					block, err := aes.NewCipher(sptr.ConnWriteAES)
					if err != nil {
						loggerv1.Get().Debug("[NetSystem::Push] #%d: push msg failed on cipher (err: %v)", e, err)
						return
					}

					aead, err := cipher.NewGCM(block)
					if err != nil {
						loggerv1.Get().Debug("[NetSystem::Push] #%d: push msg failed on gcm (err: %v)", e, err)
						return
					}

					aad := [4]byte{}
					binary.LittleEndian.PutUint32(aad[:], nmsg.Sequence())

					pl = aead.Seal(nil, nonce, netMsgEntry.Payload, aad[:])
					plen = len(pl)

					if tot := plen + limbov1.NETMSG_HEADER_SIZE; tot > total {
						sptr.WriteBuffer = slices.Grow(sptr.WriteBuffer, tot-total)
						nmsg = (*limbov1.NetMessage)(&sptr.WriteBuffer[0])
						total = tot
					}
				}

				copy(sptr.WriteBuffer[limbov1.NETMSG_HEADER_SIZE:], pl)
			}

			if plen != 0 && nmsg.Flags().Contains(limbov1.FlagEncrypted) {
				plen -= limbov1.NETMSG_GCM_SIZE
			}

			builder = nmsg.Builder().PutPayloadLen(uint32(plen))

			if plen != 0 {
				builder.PutPayloadChecksum(crc32.ChecksumIEEE(pl[0:plen]))
			}

			builder.Build()

			loggerv1.Get().Debug("[NetSystem::Push] #%d writing message (id: %s) on buffer success: %s", e, netMsgEntry.Id, nmsg.String())
			sptr.WriteBufferId = netMsgEntry.Id
			sptr.WriteBufferPos = 0

			// _ = sptr.OutcumQueue.Pop()

			return
		}

		if sptr.WriteBufferPos < limbov1.NETMSG_HEADER_SIZE {
			n, err := sysutils.Write(sptr.ConnFileHandle, sptr.WriteBuffer[sptr.WriteBufferPos:limbov1.NETMSG_HEADER_SIZE])

			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK || n == 0 {
				return
			}

			if n > 0 && err == nil {
				sptr.WriteBufferPos += n
				sptr.BytesOut += uint64(n)
				loggerv1.Get().Debug("[NetSystem::Push] #%d message (id: %s) writing progress: %d of %d header bytes", e, sptr.WriteBufferId, sptr.WriteBufferPos, limbov1.NETMSG_HEADER_SIZE)
			}

			if err != nil {
				loggerv1.Get().Debug("[NetSystem::Push] #%d message (id: %s) write err: %v (count: %d; total: %d)", e, sptr.WriteBufferId, err, sptr.WriteBufferPos, limbov1.NETMSG_HEADER_SIZE)
				return
			}
		}

		if sptr.WriteBufferPos >= limbov1.NETMSG_HEADER_SIZE {
			netMsg := limbov1.NetMessageOf(sptr.WriteBuffer)

			n, err := sysutils.Write(sptr.ConnFileHandle, sptr.WriteBuffer[sptr.WriteBufferPos:netMsg.TotalLen()])

			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK || n == 0 {
				return
			}

			if n > 0 && err == nil {
				sptr.WriteBufferPos += n
				sptr.BytesOut += uint64(n)
				loggerv1.Get().Debug("[NetSystem::Push] #%d message (%s) write progress (%d:%d) bytes", e, sptr.WriteBufferId, sptr.WriteBufferPos, netMsg.TotalLen())
			}

			if err != nil {
				return
			}

			if netMsg.TotalLen() == uint32(sptr.WriteBufferPos) {
				loggerv1.Get().Debug("[NetSystem::Push] #%d message (%s) fully writen (%d:%d) bytes: %v", e, sptr.WriteBufferId, sptr.WriteBufferPos, netMsg.TotalLen(), sptr.WriteBuffer)
				sptr.SeqOutno++
				_ = sptr.OutcumQueue.Pop()
				sptr.WriteBufferId = ""
				sptr.WriteBuffer = nil
			}
		}
	})
}

func (x *NetSystem) Accept() {
	loggerv1.Get().Debug("[NetSystem::Accept] accept connections (total: %d)", limbov1.Networks().Count())

	count := 0

	for {
		count++

		if limbov1.Networks().Count() >= int(x.maxConnections) {
			loggerv1.Get().Debug("[NetSystem::Accept] accept fail (err: limit)")
			return
		}

		if count >= 64 {
			return
		}

		x.listener.SetDeadline(time.Now().Add(time.Microsecond))

		conn, err := x.listener.AcceptTCP()

		if err != nil {
			loggerv1.Get().Debug("[NetSystem::Accept] accept call failed on accept (err: %v)", err)
			return
		}

		var fileDescriptor uintptr
		if rawConn, err := conn.SyscallConn(); true {
			if err != nil {
				conn.Close()
				continue
			}

			rawConn.Control(func(fd uintptr) {
				sysutils.SetNonblock(fd, true)
				fileDescriptor = fd
			})
		}

		if err := conn.SetNoDelay(true); err != nil {
			conn.Close()
			continue
		}

		conn.SetKeepAlive(true)
		conn.SetKeepAlivePeriod(60 * time.Second)

		connId := limbov1.NetConn(0)
		if err := limbov1.Networks().Connect(conn, TcpConnType, &connId); errnov1.FAIL(err) {
			conn.Close()
			continue
		}

		ent := limbov1.Entities().Create()

		if ptr := limbov1.NewComponentB[NetClient](ent, NetClientType); ptr != nil {
			var buff [8]byte
			rand.Read(buff[:])

			ptr.Id = binary.LittleEndian.Uint64(buff[:])
			ptr.ConnId = connId
			ptr.ConnFileHandle = fileDescriptor

			ptr.ReadBuffer = make([]byte, limbov1.NETMSG_HEADER_SIZE)[:limbov1.NETMSG_HEADER_SIZE]

			ptr.WriteBufferId = ""

			ptr.OutcumQueue = queuev1.QueueOf[*NetMessageEntry](8)

			ptr.CreatedAt = time.Now().Unix()
			ptr.TouchedAt = time.Now().Unix()

			loggerv1.Get().Debug("[NetSystem::Accept] #%d connected (sid: %d)", ent, ptr.Id)

			continue
		}

		limbov1.Entities().Destroy(ent)
		limbov1.Networks().Destroy(connId)

		conn.Close()
	}
}

func (x *NetSystem) SendMessage(typ limbov1.NetMessageType, payload proto.Message, data []byte, e limbov1.Entity) bool {
	var buff []byte
	var err error

	if payload != nil {
		if buff, err = proto.Marshal(payload); err != nil {
			loggerv1.Get().Debug("[NetSystem::SendMessage] msg (type: %d) marshaling failed (err: %v)", typ, err)
			return false
		}
	}

	return x.sendMessage(typ, buff, data, e)
}

func (x *NetSystem) sendMessage(typ limbov1.NetMessageType, payload []byte, data []byte, e limbov1.Entity) bool {
	if !limbov1.Entities().IsAlive(e) {
		loggerv1.Get().Debug("[NetSystem::sendMessage] msg send failed (err: %d is died)", e)
		return false
	}

	if sptr := limbov1.ComponentPtr[NetClient](e, NetClientType); sptr != nil {
		sptr.OutcumQueue.PushBack(&NetMessageEntry{
			Id:        uuid.NewString(),
			Typ:       typ,
			Payload:   payload,
			Any:       data,
			CreatedAt: time.Now(),
		})

		loggerv1.Get().Debug("[NetSystem::sendMessage] message (%d) will send async to #%d: %v", typ, e, payload)

		return true
	}

	return false
}

func (x *NetSystem) SendMessageListed(typ limbov1.NetMessageType, payload proto.Message, data []byte, recipients []limbov1.Entity) bool {
	var buff []byte
	var err error

	if payload != nil {
		if buff, err = proto.Marshal(payload); err != nil {
			loggerv1.Get().Debug("[NetSystem::SendMessageListed] msg (type: %d) marshaling failed (err: %v)", typ, err)
			return false
		}
	}

	for i := range len(recipients) {
		x.sendMessage(typ, buff, data, recipients[i])
	}

	return true
}
