package soupev1

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"io"
	"slices"
	"syscall"
	"time"
	"unsafe"

	limbov1 "github.com/limboware/limbo"
	loggerv1 "github.com/limboware/pkg/logger/v1"
	pbsoupev1 "github.com/limboware/pkg/proto/soupe/v1"
	sysutils "github.com/limboware/soupe/utils/syscalls"
	errnov1 "github.com/rejchev/errno"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/hkdf"
	"google.golang.org/protobuf/proto"
)

type NetInputSystem struct {
	sessionsMax              int // per ip
	sessionsTimeOut          time.Duration
	sessionsHandshakeTimeOut time.Duration

	desyncMax           uint8
	rateLimit           int // per sec
	rotationKeyInternal time.Duration
	replayWindow        int
	sequenceJumpMax     uint32
}

func netInputSystem(buff *limbov1.System) {
	v := new(NetInputSystem)

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
func (x *NetInputSystem) Activate() bool {
	return true
}

// Deactivate implements [limbov1.ISystem].
func (x *NetInputSystem) Deactivate() {

}

// Load implements [limbov1.ISystem].
func (x *NetInputSystem) Load() errnov1.Code {
	x.sessionsTimeOut = 60 * time.Second
	x.sequenceJumpMax = 10000
	x.desyncMax = 3

	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(Conn)) }, &ConnType) {
		return errnov1.ECALL
	}

	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(Client)) }, &ClientType) {
		return errnov1.ECALL
	}

	limbov1.Events().Subscribe("world.loaded", x.onAllLoaded)
	limbov1.Events().Subscribe("entity.destroy", x.onEntityDestroy)

	return errnov1.OK
}

func (x *NetInputSystem) onAllLoaded(string, any) {
	limbov1.NetworkHandlers().Register(TcpConnType, limbov1.NetMessageType(pbsoupev1.MsgType_Hello), x.helloMsgHandler)
}

func (x *NetInputSystem) helloMsgHandler(e limbov1.Entity, data *byte, count uint32) errnov1.Code {
	netMessage := (*limbov1.NetMessage)(data)

	if netMessage.PayloadLen() < 1 {
		return errnov1.EINVAL
	}

	connPtr := limbov1.ComponentPtr[Conn](e, ConnType)

	if connPtr == nil || connPtr.HandshakedAt != 0 {
		return errnov1.ECALL
	}

	payload := pbsoupev1.MsgHello{}
	if err := proto.Unmarshal(netMessage.BytesB(limbov1.NETMSG_HEADER_SIZE, int(netMessage.PayloadLen())), &payload); err != nil {
		return errnov1.EINVAL
	}

	if len(payload.Pubkey) < 32 || len(payload.Pubkey) > 36 {
		loggerv1.Get().Err("[NetInputSystem::helloMsgHandler] #%d pubkey check failed (len: %d)", e, len(payload.Pubkey))
		return errnov1.EINVAL
	}

	serverPrivKey := make([]byte, 32)

	rand.Read(serverPrivKey)

	serverPrivKey[0] &= 248
	serverPrivKey[31] &= 127
	serverPrivKey[31] |= 64

	serverPubKey, _ := curve25519.X25519(serverPrivKey, curve25519.Basepoint)

	sharedSecret, err := curve25519.X25519(serverPrivKey, payload.Pubkey)

	if err != nil {
		loggerv1.Get().Debug("client %d shared fail (err: %v): %v", e, err, payload.Pubkey)
		return errnov1.EINVAL
	}

	connPtr.Sharedkey = sharedSecret
	connPtr.ServerPubKey = serverPubKey
	connPtr.ServerPrivKey = serverPrivKey
	connPtr.PubKey = payload.Pubkey

	if err := x.deriveKeys(e, connPtr); errnov1.FAIL(err) {
		return err
	}

	connPtr.SeqInno = 0
	connPtr.SeqOutno = 0

	slotID := uint8(0)
	if !Slots().First(&slotID, func(u uint8) bool {
		return Slots().IsAvailable(u) && !Tickets().IsExpired(Slots().Ticket(u))
	}) {
		return errnov1.EINVAL
	}

	if !SendNetMessageD(limbov1.NetMessageType(pbsoupev1.MsgType_Welcome), &pbsoupev1.MsgHello_Response{
		Pubkey: serverPubKey,
	}, e) {
		return errnov1.ECALL
	}

	Tickets().Remove(Slots().Ticket(slotID))

	clientPtr := (*Client)(nil)
	if clientPtr = limbov1.ComponentPtr[Client](e, ClientType); clientPtr == nil {
		clientPtr = limbov1.NewComponentB[Client](e, ClientType)
	}

	clientPtr.Id = Slots().Owner(slotID)
	clientPtr.SlotID = slotID
	clientPtr.ConnectedAt = time.Now().Unix()

	connPtr.HandshakedAt = time.Now().Unix()

	return errnov1.ECALL
}

func (x *NetInputSystem) deriveKeys(e limbov1.Entity, ptr *Conn) errnov1.Code {
	sid := make([]byte, 8)[:8]
	binary.BigEndian.PutUint64(sid, ptr.SessionID)

	salt := fmt.Appendf(nil, "%s_skv%d", string(sid), limbov1.NETMSG_VERSION)

	hkdf := hkdf.New(sha512.New, ptr.Sharedkey, salt, nil)

	ptr.ReadAES = make([]byte, 32)[:32]
	if _, err := io.ReadFull(hkdf, ptr.ReadAES); err != nil {
		return errnov1.ECALL
	}

	ptr.WriteAES = make([]byte, 32)[:32]
	if _, err := io.ReadFull(hkdf, ptr.WriteAES); err != nil {
		return errnov1.ECALL
	}

	ptr.ReadNonce = make([]byte, 12)[:12]
	if _, err := io.ReadFull(hkdf, ptr.ReadNonce); err != nil {
		return errnov1.ECALL
	}

	ptr.WriteNonce = make([]byte, 12)[:12]
	if _, err := io.ReadFull(hkdf, ptr.WriteNonce); err != nil {
		return errnov1.ECALL
	}

	ptr.RotateKey = make([]byte, 32)[:32]
	if _, err := io.ReadFull(hkdf, ptr.RotateKey); err != nil {
		return errnov1.ECALL
	}

	return errnov1.OK
}

func (x *NetInputSystem) onEntityDestroy(_ string, data any) {
	if e, ok := data.(limbov1.Entity); ok {

		if ptr := limbov1.ComponentPtr[Conn](e, ConnType); ptr != nil {
			if limbov1.Networks().IsAlive(ptr.Id) {
				limbov1.Networks().Destroy(ptr.Id)
			}

			limbov1.DestroyComponent(e, ConnType)
		}

		if ptr := limbov1.ComponentPtr[Client](e, ClientType); ptr != nil {
			limbov1.DestroyComponent(e, ClientType)
		}
	}
}

// Unload implements [limbov1.ISystem].
func (x *NetInputSystem) Unload() {
}

// Update implements [limbov1.ISystem].
func (x *NetInputSystem) Update(dt time.Duration) {
	x.recv()
}

func (x *NetInputSystem) recv() {
	limbov1.Components().IterateB(ConnType, func(e limbov1.Entity, p unsafe.Pointer) {
		if limbov1.ContainsComponent(e, DestroyType) {
			return
		}

		connPtr := (*Conn)(p)

		connFD := uintptr(0)
		limbov1.Networks().Handle(connPtr.Id, &connFD)

		netMsgPtr := (*limbov1.NetMessage)(&connPtr.ReadBuffer[0])

		total := limbov1.NETMSG_HEADER_SIZE

		if connPtr.ReadBufferPos >= limbov1.NETMSG_HEADER_SIZE {
			total = int(netMsgPtr.TotalLen())
		}

		if connPtr.ReadBufferPos != total {
			n, err := sysutils.Read(connFD, connPtr.ReadBuffer[connPtr.ReadBufferPos:total])

			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK || n <= 0 {
				return
			}

			if err != nil {
				loggerv1.Get().Debug("[NetInputSystem::recv] #%d message read (bts %d; total: %d) err: %v", e, connPtr.ReadBufferPos, limbov1.NETMSG_HEADER_SIZE, err)
				return
			}

			connPtr.ReadBufferPos += n
			connPtr.BytesIn += uint64(n)

			if connPtr.ReadBufferPos == limbov1.NETMSG_HEADER_SIZE {
				if !x.isHeaderPotentialValid(e, connPtr, netMsgPtr.Closer(), limbov1.NETMSG_HEADER_SIZE) {
					connPtr.ReadBuffer = connPtr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
					connPtr.ReadBufferPos = 0
					connPtr.ReadDesyncCounter++
					return
				}

				total = int(netMsgPtr.TotalLen())

				if cap(connPtr.ReadBuffer) < total {
					connPtr.ReadBuffer = slices.Grow(connPtr.ReadBuffer, total-cap(connPtr.ReadBuffer))
				}
			}

			if connPtr.ReadBufferPos != total {
				return
			}
		}

		if netMsgPtr.PayloadChecksum() != netMsgPtr.PayloadRealChecksum() {
			connPtr.ReadBuffer = connPtr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
			connPtr.ReadBufferPos = 0
			connPtr.ReadDesyncCounter++
			return
		}

		if netMsgPtr.Flags().Contains(limbov1.FlagEncrypted) {
			// client nonce
			nonce := make([]byte, 12)
			copy(nonce, connPtr.ReadNonce)
			binary.LittleEndian.PutUint32(nonce[8:], binary.LittleEndian.Uint32(nonce[8:])^netMsgPtr.Sequence())

			// AES-GCM decrypting
			block, err := aes.NewCipher(connPtr.ReadAES)

			if err != nil {
				connPtr.ReadBuffer = connPtr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
				connPtr.ReadBufferPos = 0
				return
			}

			aead, err := cipher.NewGCM(block)
			if err != nil {
				connPtr.ReadBuffer = connPtr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
				connPtr.ReadBufferPos = 0
				return
			}

			text, err := aead.Open(nil, nonce, connPtr.ReadBuffer[limbov1.NETMSG_HEADER_SIZE:], connPtr.ReadBuffer[:limbov1.NETMSG_HEADER_SIZE])
			if err != nil {
				connPtr.ReadBuffer = connPtr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
				connPtr.ReadBufferPos = 0
				return
			}

			if szDiv := int(total-limbov1.NETMSG_HEADER_SIZE) - len(text); szDiv < 0 {
				szDiv *= -1
				total += szDiv
				connPtr.ReadBuffer = slices.Grow(connPtr.ReadBuffer, szDiv)
			}

			copy(connPtr.ReadBuffer[limbov1.NETMSG_HEADER_SIZE:], text)
		}

		connPtr.SeqInno = netMsgPtr.Sequence()

		if x.isMessageHandlable(e, connPtr, netMsgPtr.Closer()) {
			loggerv1.Get().Debug("[NetInputSystem::recv] #%d handler conn type: %d; msg type: %d -> %d", e, connPtr.Id.Type(), netMsgPtr.Type(), limbov1.MakeNetworkHandler(
				connPtr.Id.Type(),
				netMsgPtr.Type(),
			))
			if handlerFn := limbov1.NetworkHandlers().Get(limbov1.MakeNetworkHandler(connPtr.Id.Type(), netMsgPtr.Type())); handlerFn != nil {
				if errno := handlerFn(e, netMsgPtr.Closer(), uint32(total)); errnov1.SUCCESS(errno) {
					connPtr.TouchedAt = time.Now().Unix()
				}
			}
		}

		connPtr.ReadBuffer = connPtr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
		connPtr.ReadBufferPos = 0
		connPtr.IdleDuration = 0
	})
}

func (x *NetInputSystem) isMessageHandlable(e limbov1.Entity, connPtr *Conn, buff *byte) bool {
	netMsg := (*limbov1.NetMessage)(buff)

	// only hello msg if not handshaked
	if netMsg.Type() != limbov1.NetMessageType(pbsoupev1.MsgType_Hello) && connPtr.HandshakedAt == 0 {
		return false
	}

	// only flag none handshake
	if netMsg.Type() == limbov1.NetMessageType(pbsoupev1.MsgType_Hello) && !netMsg.Flags().Is(limbov1.FlagNone) {
		return false
	}

	// only none flags or flag enc (temp)
	return netMsg.Flags().Is(limbov1.FlagNone) || netMsg.Flags().Is(limbov1.FlagEncrypted)
}

func (x *NetInputSystem) isHeaderPotentialValid(e limbov1.Entity, connPtr *Conn, buff *byte, l int) bool {
	netMsg := (*limbov1.NetMessage)(buff)

	if netMsg.Magic() != limbov1.NETMSG_MAGIC_VALUE || netMsg.Version() != limbov1.NETMSG_VERSION {
		return false
	}

	if connPtr.HandshakedAt != 0 && netMsg.SessionID() != connPtr.SessionID {
		return false
	}

	if connPtr.HandshakedAt != 0 && netMsg.KeyID() != connPtr.KeyGen {
		return false
	}

	if connPtr.HandshakedAt != 0 && netMsg.Sequence() < connPtr.SeqInno {
		return false
	}

	if seqdiv := netMsg.Sequence() - connPtr.SeqInno; seqdiv >= x.sequenceJumpMax {
		return false
	}

	if netMsg.HeaderChecksum() != netMsg.HeaderRealChecksum() {
		return false
	}

	totalLen := netMsg.PayloadLen() + limbov1.NETMSG_HEADER_SIZE

	if netMsg.Flags().Contains(limbov1.FlagEncrypted) {
		totalLen += limbov1.NETMSG_GCM_SIZE
	}

	if netMsg.PayloadLen() > limbov1.NETMSG_PAYLOAD_MAX_LENGTH || totalLen > limbov1.NETMSG_SIZE || netMsg.TotalLen() != totalLen {
		return false
	}

	return true
}

func (x *NetInputSystem) disconnect(e limbov1.Entity, reason string) {
	dptr := limbov1.NewComponentB[Destroy](e, DestroyType)

	if dptr != nil && reason != "" {
		dptr.Reason = reason
	}

	loggerv1.Get().Debug("#%d disconnected cause %s: %v", e, reason, dptr)
}
