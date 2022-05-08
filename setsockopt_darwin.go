package tcpoption

import (
	"os"
	"syscall"
	"time"
)

// netinet/tcp.h
const (
	DARWIN_TCP_KEEPIDLE  int = 0x10
	DARWIN_TCP_KEEPINTVL int = 0x101
	DARWIN_TCP_KEEPCNT   int = 0x102
	DARWIN_TCP_FASTOPEN  int = 0x105
)

// netinet/tcp_var.h
const (
	DARWIN_TCP_FASTOPEN_SERVER int = 0x01
	DARWIN_TCP_FASTOPEN_CLIENT int = 0x02
)

// sys/socket.h
const (
	DARWIN_SO_REUSEADDR int = 0x0004
	DARWIN_SO_REUSEPORT int = 0x0200
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

func getsockoptKeepAliveIdle(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPIDLE)
}

func setsockoptKeepAliveInterval(fd int, sec int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPINTVL, sec),
	)
}

func getsockoptKeepAliveInterval(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPINTVL)
}

func setsockoptKeepAliveProbes(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPCNT, count),
	)
}

func getsockoptKeepAliveProbes(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_KEEPCNT)
}

func setsockoptFastOpen(fd int, count int) error {
	//return os.NewSyscallError(
	//	"setsockopt",
	//	syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_FASTOPEN, DARWIN_TCP_FASTOPEN_SERVER),
	//)
	return nil // not support
}

func getsockoptFastOpen(fd int) (int, error) {
	//return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_FASTOPEN)
	return 0, nil // not support
}

func setsockoptFastOpenConnect(fd int, count int) error {
	//return os.NewSyscallError(
	//	"setsockopt",
	//	syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_FASTOPEN, DARWIN_TCP_FASTOPEN_CLIENT),
	//)
	return nil // not support
}

func getsockoptFastOpenConnect(fd int) (int, error) {
	//return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, DARWIN_TCP_FASTOPEN)
	return 0, nil // not support
}

func setsockoptQuickACK(fd int, onoff int) error {
	return nil // not support
}

func setsockoptDeferAccept(fd int, onoff int) error {
	return nil // not support
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
