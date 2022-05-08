package tcpoption

import (
	"context"
	"fmt"
	"net"
	"os"
	"syscall"
	"testing"
)

func setupServer(t *testing.T, fn func(fd int) error) (string, context.CancelFunc) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "[0.0.0.0]:0")
	if err != nil {
		t.Fatalf("resolv err: %+v", err)
	}
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		t.Fatalf("resolv err: %+v", err)
	}
	defer syscall.Close(fd)

	if err := fn(fd); err != nil {
		t.Fatalf("setsockopt err: %+v", err)
	}

	sock := &syscall.SockaddrInet4{Port: tcpAddr.Port}
	if err := syscall.Bind(fd, sock); err != nil {
		t.Fatalf("bind err: %+v", err)
	}

	if err := syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		t.Fatalf("listen err: %+v", err)
	}

	f := os.NewFile(uintptr(fd), "/tmp/test")
	defer f.Close()

	listener, err := net.FileListener(f)
	if err != nil {
		t.Fatalf("listen err: %+v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer listener.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			conn, err := listener.Accept()
			if err != nil {
				return
			}
			conn.Write([]byte(t.Name()))
		}
	}()
	addr := fmt.Sprintf("127.0.0.1:%d", listener.Addr().(*net.TCPAddr).Port)
	return addr, cancel
}

func TestSetFastOpen(t *testing.T) {
	t.Skip("no support")
}

func TestSetDeferAccept(t *testing.T) {
	t.Skip("no support")
}
