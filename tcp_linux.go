package tcpoption

import (
	"os"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func tcpSetLingerTimeout(fd int, d time.Duration) error {
	tval := syscall.NsecToTimeval(d.Nanoseconds())
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptTimeval(fd, syscall.IPPROTO_TCP, syscall.TCP_LINGER2, &tval),
	)
}

func tcpSetKeepAliveIdle(fd int, sec int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPIDLE, sec),
	)
}

func tcpSetKeepAliveInterval(fd int, sec int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, sec),
	)
}

func tcpSetKeepAliveProbes(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT, count),
	)
}

func tcpSetFastOpen(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, unix.TCP_FASTOPEN, onoff),
	)
}

func tcpSetFastOpenConnect(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, unix.TCP_FASTOPEN_CONNECT, onoff),
	)
}

func tcpSetQuickACK(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_QUICKACK, onoff),
	)
}

func tcpSetDeferAccept(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_DEFER_ACCEPT, onoff),
	)
}

func tcpSetReuseAddr(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, unix.SO_REUSEADDR, onoff),
	)
}

func tcpSetReusePort(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, unix.SO_REUSEPORT, onoff),
	)
}
