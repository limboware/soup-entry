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

	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(NetClient)) }, &NetClientType) {
		return errnov1.ECALL
	}

	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(Client)) }, &ClientType) {
		return errnov1.ECALL
	}

	limbov1.Events().Subscribe("world.loaded", x.onAllLoaded)
	limbov1.Events().Subscribe("entity.destroy", x.onEntityDestroy)

	return errnov1.OK
}

// OnAllLoaded implements [limbov1.ISystem].
func (x *NetInputSystem) onAllLoaded(string, any) {
	limbov1.NetworkHandlers().Register(TcpConnType, limbov1.NetMessageType(pbsoupev1.MsgType_Hello), x.helloMsgHandler)
}

func (x *NetInputSystem) helloMsgHandler(e limbov1.Entity, data *byte, count uint32) errnov1.Code {
	netMessage := (*limbov1.NetMessage)(data)

	if netMessage.PayloadLen() < 1 {
		return errnov1.EINVAL
	}

	sptr := limbov1.ComponentPtr[NetClient](e, NetClientType)

	if sptr == nil || sptr.HandshakedAt != 0 {
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

	ticketId := TicketFrom(netMessage.Reserved())

	if !Tickets().IsAlive(ticketId) {
		return errnov1.EINVAL
	}

	sptr.SteamId = 0
	if sptr.SteamId = Tickets().Requester(ticketId); sptr.SteamId == 0 {
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

	sptr.ConnSharedkey = sharedSecret
	sptr.ConnServerPubKey = serverPubKey
	sptr.ConnServerPrivKey = serverPrivKey
	sptr.ConnPubKey = payload.Pubkey

	if err := x.deriveKeys(e, sptr); errnov1.FAIL(err) {
		return err
	}

	sptr.SeqInno = 0
	sptr.SeqOutno = 0

	if SendNetMessageD(limbov1.NetMessageType(pbsoupev1.MsgType_Welcome), &pbsoupev1.MsgHello_Response{
		Pubkey: serverPubKey,
	}, e) {

		_ = Tickets().Activate(ticketId)

		sptr.TouchedAt = time.Now().Unix()
		sptr.HandshakedAt = time.Now().Unix()
		return errnov1.OK
	}

	return errnov1.ECALL
}

func (x *NetInputSystem) deriveKeys(e limbov1.Entity, ptr *NetClient) errnov1.Code {
	sid := make([]byte, 8)[:8]
	binary.BigEndian.PutUint64(sid, ptr.Id)

	salt := fmt.Appendf(nil, "%s_skv%d", string(sid), limbov1.NETMSG_VERSION)

	hkdf := hkdf.New(sha512.New, ptr.ConnSharedkey, salt, nil)

	ptr.ConnReadAES = make([]byte, 32)[:32]
	if _, err := io.ReadFull(hkdf, ptr.ConnReadAES); err != nil {
		return errnov1.ECALL
	}

	ptr.ConnWriteAES = make([]byte, 32)[:32]
	if _, err := io.ReadFull(hkdf, ptr.ConnWriteAES); err != nil {
		return errnov1.ECALL
	}

	ptr.ConnReadNonce = make([]byte, 12)[:12]
	if _, err := io.ReadFull(hkdf, ptr.ConnReadNonce); err != nil {
		return errnov1.ECALL
	}

	ptr.ConnWriteNonce = make([]byte, 12)[:12]
	if _, err := io.ReadFull(hkdf, ptr.ConnWriteNonce); err != nil {
		return errnov1.ECALL
	}

	ptr.ConnRotateKey = make([]byte, 32)[:32]
	if _, err := io.ReadFull(hkdf, ptr.ConnRotateKey); err != nil {
		return errnov1.ECALL
	}

	return errnov1.OK
}

func (x *NetInputSystem) onEntityDestroy(_ string, data any) {
	if e, ok := data.(limbov1.Entity); ok {

		if ptr := limbov1.ComponentPtr[NetClient](e, NetClientType); ptr != nil {
			if limbov1.Networks().IsAlive(ptr.ConnId) {
				limbov1.Networks().Destroy(ptr.ConnId)
			}

			limbov1.DestroyComponent(e, NetClientType)
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
	x.recv(dt)
}

func (x *NetInputSystem) recv(dt time.Duration) {
	limbov1.Components().IterateB(NetClientType, func(e limbov1.Entity, p unsafe.Pointer) {
		if limbov1.ContainsComponent(e, DestroyType) {
			return
		}

		ndptr := (*NetClient)(p)

		ndptr.IdleDuration++

		if times := time.Duration(ndptr.IdleDuration) * dt; times >= x.sessionsTimeOut || ndptr.ReadDesyncCounter >= x.desyncMax {
			// ndptr.IdleDuration = 0

			if limbov1.Networks().IsAlive(ndptr.ConnId) {
				ndptr.ConnFileHandle = 0
				limbov1.Networks().Destroy(ndptr.ConnId)
			}

			if times >= x.sessionsTimeOut+30*time.Second {
				loggerv1.Get().Debug("#%d session disconnect cause timeout (t: %d; tt: %d)", e, times, x.sessionsTimeOut+30*time.Second)
				x.disconnect(e, "timeout")
				return
			}

			if ndptr.ReadDesyncCounter >= x.desyncMax {
				x.disconnect(e, "desynced")
			}

			return
		}

		if ndptr.ConnFileHandle == 0 {
			return
		}

		if ndptr.ReadBufferPos < limbov1.NETMSG_HEADER_SIZE {

			n, err := sysutils.Read(ndptr.ConnFileHandle, ndptr.ReadBuffer[ndptr.ReadBufferPos:limbov1.NETMSG_HEADER_SIZE])

			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK || n == 0 {
				return
			}

			if n > 0 && err == nil {
				ndptr.IdleDuration = 0
				ndptr.ReadBufferPos += n
				ndptr.BytesIn += uint64(n)

				loggerv1.Get().Debug("[NetInputSystem::recv] #%d message read progress: %d (total: %d)", e, ndptr.ReadBufferPos, limbov1.NETMSG_HEADER_SIZE)

				if ndptr.ReadBufferPos >= limbov1.NETMSG_HEADER_SIZE {
					if !x.isHeaderPotentialValid(e, ndptr, &ndptr.ReadBuffer[0], limbov1.NETMSG_HEADER_SIZE) {
						ndptr.ReadBuffer = ndptr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
						ndptr.ReadBufferPos = 0
						ndptr.ReadDesyncCounter++
						return
					}

					if ndptr.ReadDesyncCounter > 0 {
						ndptr.ReadDesyncCounter = 0
					}
				}
			}

			// not interesting
			if err != nil {
				loggerv1.Get().Debug("[NetInputSystem::recv] #%d message read (bts %d; total: %d) err: %v", e, ndptr.ReadBufferPos, limbov1.NETMSG_HEADER_SIZE, err)
				return
			}
		}

		if ndptr.ReadBufferPos >= limbov1.NETMSG_HEADER_SIZE {
			nMsg := (*limbov1.NetMessage)(&ndptr.ReadBuffer[0])

			totalLen := nMsg.PayloadLen() + limbov1.NETMSG_HEADER_SIZE

			if nMsg.Flags().Contains(limbov1.FlagEncrypted) {
				totalLen += limbov1.NETMSG_GCM_SIZE
			}

			if int(nMsg.PayloadLen()) > 0 && cap(ndptr.ReadBuffer) < int(totalLen) {
				buffer := make([]byte, totalLen)[:totalLen]

				copy(buffer, ndptr.ReadBuffer)

				ndptr.ReadBuffer = buffer
			}

			szReadWindow := 4096

			szRemainingBytes := totalLen - uint32(ndptr.ReadBufferPos)

			if szRemainingBytes < uint32(szReadWindow) {
				szReadWindow = int(szRemainingBytes)
			}

			if szReadWindow > 0 {

				n, err := sysutils.Read(ndptr.ConnFileHandle, ndptr.ReadBuffer[ndptr.ReadBufferPos:ndptr.ReadBufferPos+szReadWindow])

				if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK || n == 0 {
					return
				}

				if n > 0 && err == nil {
					ndptr.ReadBufferPos += n
					ndptr.IdleDuration = 0
					ndptr.BytesIn += uint64(n)
					loggerv1.Get().Debug("[NetInputSystem::recv] #%d message read progress: %d (total: %d)", e, ndptr.ReadBufferPos, totalLen)
				}

				if err != nil {
					loggerv1.Get().Debug("[NetInputSystem::recv] #%d message read (bts %d; total: %d) err: %v", e, ndptr.ReadBufferPos, totalLen, err)
					return
				}
			}

			// прочитали сообщение целиком
			if ndptr.ReadBufferPos == int(totalLen) {
				loggerv1.Get().Debug("[NetInputSystem::recv] #%d message fully readed: %v (len: %d; total: %d)", e, ndptr.ReadBuffer, ndptr.ReadBufferPos, totalLen)

				if nMsg.PayloadChecksum() != nMsg.PayloadRealChecksum() {
					ndptr.ReadBuffer = ndptr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
					ndptr.ReadBufferPos = 0
					ndptr.ReadDesyncCounter++
					return
				}

				if nMsg.Flags().Contains(limbov1.FlagEncrypted) {
					// client nonce
					nonce := make([]byte, 12)
					copy(nonce, ndptr.ConnReadNonce)
					binary.LittleEndian.PutUint32(nonce[8:], binary.LittleEndian.Uint32(nonce[8:])^nMsg.Sequence())

					// AES-GCM decrypting
					block, err := aes.NewCipher(ndptr.ConnReadAES)

					if err != nil {
						ndptr.ReadBuffer = ndptr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
						ndptr.ReadBufferPos = 0
						return
					}

					aead, err := cipher.NewGCM(block)
					if err != nil {
						ndptr.ReadBuffer = ndptr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
						ndptr.ReadBufferPos = 0
						return
					}

					text, err := aead.Open(nil, nonce, ndptr.ReadBuffer[limbov1.NETMSG_HEADER_SIZE:totalLen], ndptr.ReadBuffer[:limbov1.NETMSG_HEADER_SIZE])
					if err != nil {
						ndptr.ReadBuffer = ndptr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
						ndptr.ReadBufferPos = 0
						return
					}

					if szDiv := int(totalLen-limbov1.NETMSG_HEADER_SIZE) - len(text); szDiv < 0 {
						szDiv *= -1
						totalLen += uint32(szDiv)
						ndptr.ReadBuffer = slices.Grow(ndptr.ReadBuffer, szDiv)
					}

					copy(ndptr.ReadBuffer[limbov1.NETMSG_HEADER_SIZE:], text)
				}

				ndptr.SeqInno = nMsg.Sequence()

				if x.isMessageHandlable(e, ndptr, nMsg.Closer()) {
					loggerv1.Get().Debug("[NetInputSystem::recv] #%d handler conn type: %d; msg type: %d -> %d", e, ndptr.ConnId.Type(), nMsg.Type(), limbov1.MakeNetworkHandler(
						ndptr.ConnId.Type(),
						nMsg.Type(),
					))
					if handlerFn := limbov1.NetworkHandlers().Get(limbov1.MakeNetworkHandler(
						ndptr.ConnId.Type(),
						nMsg.Type(),
					)); handlerFn != nil {
						handlerFn(e, nMsg.Closer(), totalLen)
					}
				}

				ndptr.ReadBuffer = ndptr.ReadBuffer[:][:limbov1.NETMSG_HEADER_SIZE]
				ndptr.ReadBufferPos = 0
				ndptr.IdleDuration = 0
			}
		}
	})
}

func (x *NetInputSystem) isMessageHandlable(e limbov1.Entity, sptr *NetClient, buff *byte) bool {
	netMsg := (*limbov1.NetMessage)(buff)

	// only hello msg if not handshaked
	if netMsg.Type() != limbov1.NetMessageType(pbsoupev1.MsgType_Hello) && sptr.HandshakedAt == 0 {
		return false
	}

	// only flag none handshake
	if netMsg.Type() == limbov1.NetMessageType(pbsoupev1.MsgType_Hello) && !netMsg.Flags().Is(limbov1.FlagNone) {
		return false
	}

	// only none flags or flag enc (temp)
	return netMsg.Flags().Is(limbov1.FlagNone) || netMsg.Flags().Is(limbov1.FlagEncrypted)
}

func (x *NetInputSystem) isHeaderPotentialValid(e limbov1.Entity, sptr *NetClient, buff *byte, l int) bool {
	netMsg := (*limbov1.NetMessage)(buff)

	if netMsg.Magic() != limbov1.NETMSG_MAGIC_VALUE || netMsg.Version() != limbov1.NETMSG_VERSION {
		return false
	}

	if sptr.HandshakedAt != 0 && netMsg.SessionID() != sptr.Id {
		return false
	}

	if sptr.HandshakedAt != 0 && netMsg.KeyID() != sptr.ConnKeyGen {
		return false
	}

	if sptr.HandshakedAt != 0 && netMsg.Sequence() < sptr.SeqInno {
		return false
	}

	if seqdiv := netMsg.Sequence() - sptr.SeqInno; seqdiv >= x.sequenceJumpMax {
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
