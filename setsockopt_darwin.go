package tcpoption

import (
	"os"
	"syscall"
	"time"
)

// netinet/tcp.h
const (
	DARWIN_TCP_KEEPIDLE  = 0x10
	DARWIN_TCP_KEEPINTVL = 0x101
	DARWIN_TCP_KEEPCNT   = 0x102
	DARWIN_TCP_FASTOPEN  = 0x105
)

// sys/socket.h
const (
	DARWIN_SO_REUSEADDR = 0x0004
	DARWIN_SO_REUSEPORT = 0x0200
)

func setsockoptLingerTimeout(fd int, d time.Duration) error {
	return nil // no option by darwin
}

func setsockoptKeepAliveIdle(fd int, sec int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPIDLE, sec),
	)
}

func setsockoptKeepAliveInterval(fd int, sec int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPINTVL, sec),
	)
}

func setsockoptKeepAliveProbes(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPCNT, count),
	)
}

func setsockoptFastOpen(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_FASTOPEN, onoff),
	)
}

func setsockoptFastOpenConnect(fd int, onoff int) error {
	return nil // no option by darwin
}

func setsockoptQuickACK(fd int, onoff int) error {
	return nil // no option by darwin
}

func setsockoptDeferAccept(fd int, onoff int) error {
	return nil // no option by darwin
}

func setsockoptReuseAddr(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, DARWIN_SO_REUSEADDR, onoff),
	)
}

func setsockoptReusePort(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, DARWIN_SO_REUSEPORT, onoff),
	)
}
