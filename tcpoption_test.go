package tcpoption

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestGetFd(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "[0.0.0.0]:0")
	if err != nil {
		t.Fatalf("resolv err: %+v", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
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

	tcpConn := conn.(*net.TCPConn)
	f, err := tcpConn.File()
	if err != nil {
		t.Fatalf("client fd open: %+v", err)
	}
	defer f.Close()

	fileFd := int(f.Fd())
	if err := getFd(tcpConn, func(fd int) error {
		t.Logf("fileFd: %d, rawControlFd: %d", fileFd, fd)
		return nil
	}); err != nil {
		t.Errorf("must no error: %+v", err)
	}

	listener.Close()
	wg.Wait()
}
