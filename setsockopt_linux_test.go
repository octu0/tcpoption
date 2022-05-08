package tcpoption

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

type testingWrap interface {
	Fatalf(fmt string, v ...interface{})
	Logf(fmt string, v ...interface{})
	Name() string
}

func setupServer(t testingWrap, fn func(fd int) error) (string, context.CancelFunc, *sync.WaitGroup) {
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

	open := int64(0)
	send := func(wg *sync.WaitGroup, conn net.Conn) {
		defer wg.Done()
		defer atomic.AddInt64(&open, -1)

		conn.Write([]byte(t.Name()))
		conn.Close()
	}
	ctx, cancel := context.WithCancel(context.Background())
	mainWG := new(sync.WaitGroup)
	mainWG.Add(1)
	go func() {
		<-ctx.Done()
		listener.Close()
	}()
	go func() {
		i := 0
		wg := new(sync.WaitGroup)
		defer func() {
			wg.Wait()
			mainWG.Done()
			//t.Logf("%s total open %d", t.Name(), i)
		}()

		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}

			atomic.AddInt64(&open, 1)
			wg.Add(1)
			go send(wg, conn)

			// stats
			i += 1
			if (i % 100) == 0 {
				//t.Logf("%s open %d/%d", t.Name(), atomic.LoadInt64(&open), i)
			}
		}
	}()
	addr := fmt.Sprintf("127.0.0.1:%d", listener.Addr().(*net.TCPAddr).Port)
	return addr, cancel, mainWG
}

func benchmarkApp(b *testing.B, setupFD func(fd int) error) {
	pool := new(sync.Pool)
	getTimer := func(dur time.Duration) *time.Timer {
		v := pool.Get()
		if v == nil {
			return time.NewTimer(dur)
		}
		tm := v.(*time.Timer)
		tm.Reset(dur)
		return tm
	}
	putTimer := func(tm *time.Timer) {
		if tm.Stop() != true {
			select {
			case <-tm.C:
				// drain
			default:
			}
		}
		pool.Put(tm)
	}

	rand.Seed(time.Now().UnixNano())
	slowClient := func(tb *testing.B, wg *sync.WaitGroup, addr string) {
		defer wg.Done()

		rnd := time.Duration(rand.Intn(50)+1) * time.Millisecond
		base := 100 * time.Millisecond
		tm := getTimer(base + rnd)
		defer putTimer(tm)

		conn, err := net.Dial("tcp", addr)
		if err != nil {
			tb.Fatalf("client: %+v", err)
		}
		defer conn.Close()

		<-tm.C

		conn.Write([]byte("PING"))
	}

	count := 1000
	b.Run("disable", func(tb *testing.B) {
		addr, done, svr := setupServer(tb, func(fd int) error {
			return nil
		})
		tb.ResetTimer()

		c := new(sync.WaitGroup)
		c.Add(count)
		for i := 0; i < count; i += 1 {
			go slowClient(tb, c, addr)
		}
		c.Wait()
		done()
		svr.Wait()
	})
	b.Run("enable", func(tb *testing.B) {
		addr, done, svr := setupServer(tb, setupFD)
		tb.ResetTimer()

		c := new(sync.WaitGroup)
		c.Add(count)
		for i := 0; i < count; i += 1 {
			go slowClient(tb, c, addr)
		}
		c.Wait()
		done()
		svr.Wait()
	})
}

func BenchmarkFastOpen(b *testing.B) {
	benchmarkApp(b, func(fd int) error {
		return setsockoptFastOpen(fd, 16*1024)
	})
}

func BenchmarkDeferAccept(b *testing.B) {
	benchmarkApp(b, func(fd int) error {
		return setsockoptDeferAccept(fd, 1)
	})
}

func TestSetFastOpen(t *testing.T) {
	addr, done, svr := setupServer(t, func(fd int) error {
		if err := setsockoptFastOpen(fd, 4*1024); err != nil {
			return err
		}

		value, err := getsockoptFastOpen(fd)
		if err != nil {
			return err
		}
		if value != (4 * 1024) {
			t.Errorf("fast open value:%d", value)
		}

		if onoff, err := getsockoptDeferAccept(fd); err != nil {
			return err
		} else {
			if onoff != 0 {
				t.Errorf("no value is set, so it should default value")
			}
		}

		return nil
	})

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("client open err: %+v", err)
	}
	conn.Write([]byte("PING"))
	conn.Close()

	done()
	svr.Wait()
}

func TestSetDeferAccept(t *testing.T) {
	addr, done, svr := setupServer(t, func(fd int) error {
		if err := setsockoptDeferAccept(fd, 1); err != nil {
			return err
		}

		onoff, err := getsockoptDeferAccept(fd)
		if err != nil {
			return err
		}
		if onoff != 1 {
			t.Errorf("enable defer_accept flag: %d", onoff)
		}

		return nil
	})

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("client open err: %+v", err)
	}
	defer conn.Close()
	conn.Write([]byte("PING"))

	done()
	svr.Wait()
}
