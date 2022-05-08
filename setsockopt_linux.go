package tcpoption

import (
	"os"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

func setsockoptLingerTimeout(fd int, d time.Duration) error {
	tval := syscall.NsecToTimeval(d.Nanoseconds())
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptTimeval(fd, syscall.IPPROTO_TCP, syscall.TCP_LINGER2, &tval),
	)
}

func setsockoptKeepAliveIdle(fd int, sec int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPIDLE, sec),
	)
}

func getsockoptKeepAliveIdle(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPIDLE)
}

func setsockoptKeepAliveInterval(fd int, sec int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, sec),
	)
}

func getsockoptKeepAliveInterval(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL)
}

func setsockoptKeepAliveProbes(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT, count),
	)
}

func getsockoptKeepAliveProbes(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT)
}

func setsockoptFastOpen(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, unix.TCP_FASTOPEN, count),
	)
}

func getsockoptFastOpen(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, unix.TCP_FASTOPEN)
}

func setsockoptFastOpenConnect(fd int, count int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, unix.TCP_FASTOPEN_CONNECT, count),
	)
}

func setsockoptQuickACK(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_QUICKACK, onoff),
	)
}

func setsockoptDeferAccept(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_DEFER_ACCEPT, onoff),
	)
}

func getsockoptDeferAccept(fd int) (int, error) {
	return syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, unix.TCP_DEFER_ACCEPT)
}

func setsockoptReuseAddr(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, unix.SO_REUSEADDR, onoff),
	)
}

func setsockoptReusePort(fd int, onoff int) error {
	return os.NewSyscallError(
		"setsockopt",
		syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, unix.SO_REUSEPORT, onoff),
	)
}
