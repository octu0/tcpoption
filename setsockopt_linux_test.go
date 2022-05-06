package tcpoption

import (
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"sync"
	"syscall"
	"testing"
)

func TestSetFastOpen(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "[0.0.0.0]:0")
	if err != nil {
		t.Fatalf("resolv err: %+v", err)
	}
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP)
	if err != nil {
		t.Fatalf("resolv err: %+v", err)
	}
	defer syscall.Close(fd)
	sock := &syscall.SockaddrInet4{Port: tcpAddr.Port}

	if err := tcpSetFastOpen(fd, 4*1024); err != nil {
		t.Errorf("must no error: %+v", err)
	}

	v, err := syscall.GetsockoptInt(fd, syscall.IPPROTO_TCP, unix.TCP_FASTOPEN)
	if err != nil {
		t.Errorf("must no error: %+v", err)
	}

	if v != (4 * 1024) {
		t.Errorf("setsockopt val:%d", v)
	}

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
	defer listener.Close()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(_wg *sync.WaitGroup) {
		defer _wg.Done()

		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			conn.Write([]byte("test"))
		}
	}(wg)

	conn, err := net.Dial("tcp", fmt.Sprintf(
		"127.0.0.1:%d",
		listener.Addr().(*net.TCPAddr).Port,
	))
	if err != nil {
		t.Fatalf("client open err: %+v", err)
	}
	defer conn.Close()

	listener.Close()
	wg.Wait()
}
