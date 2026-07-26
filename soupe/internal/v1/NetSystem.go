package soupev1

import (
	"net"
	"syscall"
	"time"

	configsv2 "github.com/limboware/pkg/configs/v2"
	limbov1 "github.com/limboware/pkg/limbo/v1"
	sysutils "github.com/limboware/soupe/utils/syscalls"
	errnov1 "github.com/rejchev/errno"
)

var _ limbov1.ISystem = (*NetSystem)(nil)

type NetSystem struct {
	listener *net.TCPListener

	maxConnections uint32
}

// Activate implements [limbov1.ISystem].
func (x *NetSystem) Activate() bool {
	return true
}

// Deactivate implements [limbov1.ISystem].
func (x *NetSystem) Deactivate() {}

// Load implements [limbov1.ISystem].
func (x *NetSystem) Load() errnov1.Code {
	x.maxConnections = 256

	listener, err := net.Listen("tcp", ":"+configsv2.Get().ValueA("port", "27720"))

	if err != nil {
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
			x.listener.Close()
			return errnov1.ECALL
		}

		rawConn.Control(func(fd uintptr) {
			sysutils.SetNonblock(fd, true)
		})
	}

	return errnov1.OK
}

// OnAllLoaded implements [limbov1.ISystem].
func (x *NetSystem) OnAllLoaded() {

}

// Unload implements [limbov1.ISystem].
func (x *NetSystem) Unload() {
	if x.listener != nil {
		x.listener.Close()
	}
}

// Update implements [limbov1.ISystem].
func (x *NetSystem) Update(dt time.Duration) {
	x.acceptConnections()
}

func (x *NetSystem) acceptConnections() {
	for {
		if limbov1.Networks().Count() >= int(x.maxConnections) {
			return
		}

		conn, err := x.listener.AcceptTCP()

		if err != nil {
			if isEAGAIN(err) {
				return
			}

			continue
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

		connId := limbov1.ConnectionId(0)
		if err := limbov1.Networks().ConnectB(conn, limbov1.TCPConnProto.Closer(), &connId); errnov1.FAIL(err) {
			conn.Close()
			continue
		}

		ent := limbov1.Entities().Create()

		if ptr := limbov1.NewComponentB[NetConn](ent, NetConnType); ptr != nil {
			if ndptr := limbov1.NewComponentB[NetDraftConn](ent, NetDraftConnType); ndptr != nil {
				ptr.ConnId = connId
				ptr.FileHandle = fileDescriptor
				ptr.CreatedAt = time.Now().Unix()

				ndptr.ReadBuffer = make([]byte, limbov1.NETMSG_HEADER_SIZE)[:limbov1.NETMSG_HEADER_SIZE]
				ndptr.CreatedAt = time.Now().Unix()

				return
			}

			limbov1.Components().Destroy(limbov1.ComponentFor(ent, NetConnType))
		}

		limbov1.Entities().Destroy(ent)
		limbov1.Networks().Destroy(connId)

		conn.Close()
	}
}

func isEAGAIN(err error) bool {
	if err == nil {
		return false
	}
	if opErr, ok := err.(*net.OpError); ok {
		if opErr.Temporary() {
			return true
		}
		return opErr.Err == syscall.EAGAIN || opErr.Err == syscall.EWOULDBLOCK
	}
	return false
}
