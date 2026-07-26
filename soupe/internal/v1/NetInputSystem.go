package soupev1

import (
	"time"
	"unsafe"

	limbov1 "github.com/limboware/pkg/limbo/v1"
	sysutils "github.com/limboware/soupe/utils/syscalls"
	errnov1 "github.com/rejchev/errno"
)

var _ limbov1.ISystem = (*NetInputSystem)(nil)

type NetInputSystem struct {
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

	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(WebsocketConn)) }, &WebsocketConnType) {
		return errnov1.ECALL
	}

	if !limbov1.GetWorld().CreateCompotype(func() unsafe.Pointer { return unsafe.Pointer(new(TCPConn)) }, &TCPConnType) {
		return errnov1.ECALL
	}

	return errnov1.OK
}

// OnAllLoaded implements [limbov1.ISystem].
func (x *NetInputSystem) OnAllLoaded() {
	limbov1.Events().Subscribe("entity.destroy", x.onEntityDestroy)
}

func (x *NetInputSystem) onEntityDestroy(_ string, data any) {
	if e, ok := data.(limbov1.Entity); ok {
		if ptr := limbov1.ComponentPtr[WebsocketConn](e, WebsocketConnType); ptr != nil {
			limbov1.Networks().Connection(ptr.Id).Close()
			limbov1.Networks().Destroy(ptr.Id)
			limbov1.DestroyComponent(e, WebsocketConnType)
		}

		if ptr := limbov1.ComponentPtr[TCPConn](e, TCPConnType); ptr != nil {
			limbov1.Networks().Connection(ptr.Id).Close()
			limbov1.Networks().Destroy(ptr.Id)
			limbov1.DestroyComponent(e, TCPConnType)
		}
	}
}

// Unload implements [limbov1.ISystem].
func (x *NetInputSystem) Unload() {
}

// Update implements [limbov1.ISystem].
func (x *NetInputSystem) Update(dt time.Duration) {
	x.draftRecv(dt)

	x.clientsRecv(dt)
}

func (x *NetInputSystem) draftRecv(dt time.Duration) {
	limbov1.Components().IterateB(NetDraftConnType, func(e limbov1.Entity, p unsafe.Pointer) {
		ndptr := (*NetDraftConn)(p)

		ndptr.SleepTicks++

		ncptr := (*NetConn)(limbov1.ComponentPtr[NetConn](e, NetConnType))

		if (time.Duration(ndptr.SleepTicks) * dt) >= time.Second {
			ndptr.SleepTicks = 0
			x.disconnect(e, "timeout")
			return
		}

		if ndptr.ReadBufferPos < limbov1.NETMSG_HEADER_SIZE {

			n, err := sysutils.Read(ncptr.FileHandle, ndptr.ReadBuffer[ndptr.ReadBufferPos:limbov1.NETMSG_HEADER_SIZE])

			if n > 0 && err == nil {
				ndptr.SleepTicks = 0
				ndptr.ReadBufferPos += n
			}

			// not interesting, just 1 second for get netclient (connect/reconnect) or goodbuy
			if n == 0 || err != nil {
				return
			}
		}

		if ndptr.ReadBufferPos >= limbov1.NETMSG_HEADER_SIZE {
			nMsg := limbov1.NetMessageOf(ndptr.ReadBuffer)

			totalLen := nMsg.PayloadLen() + limbov1.NETMSG_HEADER_SIZE

			if (nMsg.Flags() & 0x1) != 0 {
				totalLen += limbov1.NETMSG_GCM_SIZE
			}

			if nMsg.Magic() != limbov1.NETMSG_MAGIC_VALUE ||
				nMsg.HeaderChecksum() != nMsg.HeaderRealChecksum() ||
				nMsg.SessionID() != uint64(ncptr.ConnId) ||
				nMsg.Version() != limbov1.NETMSG_VERSION ||
				nMsg.PayloadLen() > limbov1.NETMSG_PAYLOAD_MAX_LENGTH ||
				totalLen > limbov1.NETMSG_SIZE ||
				nMsg.TotalLen() != totalLen {
				ndptr.ReadBuffer = make([]byte, limbov1.NETMSG_HEADER_SIZE)[:limbov1.NETMSG_HEADER_SIZE]
				ndptr.ReadBufferPos = 0
				return
			}

			if int(nMsg.PayloadLen()) > 0 && cap(ndptr.ReadBuffer) < int(totalLen) {
				buffer := make([]byte, totalLen)[:totalLen]

				copy(buffer, ndptr.ReadBuffer)

				ndptr.ReadBuffer = buffer
			}

			// количество читаемых байтов за тик
			szReadWindow := 4096

			// остаток байт
			szRemainingBytes := totalLen - uint32(ndptr.ReadBufferPos)

			// поправка
			if szRemainingBytes < uint32(szReadWindow) {
				szReadWindow = int(szRemainingBytes)
			}

			// читаем, если есть
			if szRemainingBytes != 0 {
				n, err := sysutils.Read(ncptr.FileHandle, ndptr.ReadBuffer[ndptr.ReadBufferPos:ndptr.ReadBufferPos+szReadWindow])

				if n > 0 && err == nil {
					ndptr.ReadBufferPos += n
					ndptr.SleepTicks = 0
				}

				if n == 0 || err == nil {
					return
				}
			}

			// прочитали сообщение целиком
			if ndptr.ReadBufferPos == int(totalLen) {

				// TODO: энкрипт
				encTotal := nMsg.PayloadLen()

				if handlerFn := limbov1.NetworkHandlers().Get(limbov1.NetworkHandlerFor(
					limbov1.ConnectionProtocol(ncptr.ConnId.Protocol()),
					nMsg.Type(),
				)); handlerFn != nil {
					handlerFn(e, nMsg.Closer(), encTotal)
				}

				ndptr.ReadBuffer = make([]byte, limbov1.NETMSG_HEADER_SIZE)[:limbov1.NETMSG_HEADER_SIZE]
				ndptr.ReadBufferPos = 0
				ndptr.SleepTicks = 0
			}
		}
	})
}

func (x *NetInputSystem) clientsRecv(dt time.Duration) {
	limbov1.Components().IterateB(NetClientType, func(e limbov1.Entity, p unsafe.Pointer) {
		cptr := (*NetClient)(p)

		if cptr.ConnEnt == SERVER {
			return
		}

		cptr.SleepTicks++
	})
}

func (x *NetInputSystem) disconnect(e limbov1.Entity, reason string) {
	limbov1.NewComponentB[Destroy](e, DestroyType)
}
