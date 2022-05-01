package tcpoption

import (
	"net"
	"time"
)

func SetNoLinger(conn net.Conn, enable bool) error {
	if enable {
		if c, ok := conn.(*net.TCPConn); ok {
			return tcpSetLinger(c, 0)
		}
	}
	return nil
}

func SetLinger(conn net.Conn, d time.Duration) error {
	if d == 0 {
		return SetNoLinger(conn, true)
	}
	if c, ok := conn.(*net.TCPConn); ok {
		return tcpSetLinger(c, IntSecond(d))
	}
	return nil
}

func SetLingerTimeout(conn net.Conn, d time.Duration) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return tcpSetLingerTimeout(fd, d)
		})
	}
	return nil
}

func SetReadBuffer(conn net.Conn, bytes int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return tcpSetReadBuffer(c, bytes)
	}
	return nil
}

func SetWriteBuffer(conn net.Conn, bytes int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return tcpSetWriteBuffer(c, bytes)
	}
	return nil
}

func SetNoDelay(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return tcpSetNoDelay(c, enable)
	}
	return nil
}

func KeepAlive(conn net.Conn, enable bool, idle, interval time.Duration, probes int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := tcpSetKeepAlive(c, enable); err != nil {
			return err
		}
		if enable != true {
			return nil
		}
		return getFd(c, func(fd int) error {
			if err := tcpSetKeepAliveIdle(fd, IntSecond(idle)); err != nil {
				return err
			}
			if err := tcpSetKeepAliveInterval(fd, IntSecond(interval)); err != nil {
				return err
			}
			return tcpSetKeepAliveProbes(fd, probes)
		})
	}
	return nil
}

func SetKeepAlive(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return tcpSetKeepAlive(c, enable)
	}
	return nil
}

func SetKeepAliveTime(conn net.Conn, d time.Duration) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := tcpSetKeepAlive(c, true); err != nil {
			return err
		}
		return getFd(c, func(fd int) error {
			return tcpSetKeepAliveIdle(fd, IntSecond(d))
		})
	}
	return nil
}

func SetKeepAliveInterval(conn net.Conn, d time.Duration) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := tcpSetKeepAlive(c, true); err != nil {
			return err
		}
		return getFd(c, func(fd int) error {
			return tcpSetKeepAliveInterval(fd, IntSecond(d))
		})
	}
	return nil
}

func SetKeepAliveProbes(conn net.Conn, count int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := tcpSetKeepAlive(c, true); err != nil {
			return err
		}
		return getFd(c, func(fd int) error {
			return tcpSetKeepAliveProbes(fd, count)
		})
	}
	return nil
}

func SetFastOpen(conn net.Conn, count int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return tcpSetFastOpen(fd, count)
		})
	}
	return nil
}

func SetFastOpenConnect(conn net.Conn, count int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return tcpSetFastOpenConnect(fd, count)
		})
	}
	return nil
}

func SetQuickACK(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return tcpSetQuickACK(fd, IntBool(enable))
		})
	}
	return nil
}

func SetDeferAccept(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return tcpSetDeferAccept(fd, IntBool(enable))
		})
	}
	return nil
}

func SetReuseAddr(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return tcpSetReuseAddr(fd, IntBool(enable))
		})
	}
	return nil
}

func SetReusePort(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return tcpSetReusePort(fd, IntBool(enable))
		})
	}
	return nil
}

type Config struct {
	NoLinger          bool
	LingerTimeout     time.Duration
	ReadBuffer        int
	WriteBuffer       int
	EnableNoDelay     bool
	EnableKeepAlive   bool
	KeepAliveTime     time.Duration
	KeepAliveInterval time.Duration
	KeepAliveProbes   int
	FastOpen          int
	FastOpenConnect   int
	EnableQuickACK    bool
	EnableDeferAccept bool
	EnableReuseAddr   bool
	EnableReusePort   bool
}

func Set(conn net.Conn, cfg Config) error {
	c, ok := conn.(*net.TCPConn)
	if ok != true {
		return nil
	}
	return getFd(c, func(fd int) error {
		if cfg.NoLinger {
			if err := tcpSetLinger(c, 0); err != nil {
				return err
			}
		}
		if 0 < cfg.LingerTimeout {
			if err := tcpSetLingerTimeout(fd, cfg.LingerTimeout); err != nil {
				return err
			}
		}
		if 0 < cfg.ReadBuffer {
			if err := tcpSetReadBuffer(c, cfg.ReadBuffer); err != nil {
				return err
			}
		}
		if 0 < cfg.WriteBuffer {
			if err := tcpSetWriteBuffer(c, cfg.WriteBuffer); err != nil {
				return err
			}
		}
		if err := tcpSetNoDelay(c, cfg.EnableNoDelay); err != nil {
			return err
		}
		if err := tcpSetKeepAlive(c, cfg.EnableKeepAlive); err != nil {
			return err
		}
		if cfg.EnableKeepAlive {
			if 0 < cfg.KeepAliveTime {
				if err := tcpSetKeepAliveIdle(fd, IntSecond(cfg.KeepAliveTime)); err != nil {
					return err
				}
			}
			if 0 < cfg.KeepAliveInterval {
				if err := tcpSetKeepAliveInterval(fd, IntSecond(cfg.KeepAliveInterval)); err != nil {
					return err
				}
			}
			if 0 < cfg.KeepAliveProbes {
				if err := tcpSetKeepAliveProbes(fd, cfg.KeepAliveProbes); err != nil {
					return err
				}
			}
		}
		if 0 < cfg.FastOpen {
			if err := tcpSetFastOpen(fd, cfg.FastOpen); err != nil {
				return err
			}
		}
		if 0 < cfg.FastOpenConnect {
			if err := tcpSetFastOpenConnect(fd, cfg.FastOpenConnect); err != nil {
				return err
			}
		}
		if err := tcpSetQuickACK(fd, IntBool(cfg.EnableQuickACK)); err != nil {
			return err
		}
		if err := tcpSetDeferAccept(fd, IntBool(cfg.EnableDeferAccept)); err != nil {
			return err
		}
		if err := tcpSetReuseAddr(fd, IntBool(cfg.EnableReuseAddr)); err != nil {
			return err
		}
		if err := tcpSetReusePort(fd, IntBool(cfg.EnableReusePort)); err != nil {
			return err
		}
		return nil
	})
}

func IntSecond(d time.Duration) int {
	return int(d.Seconds())
}

func IntBool(b bool) int {
	if b {
		return 1
	}
	return 0
}

func getFd(conn *net.TCPConn, cb func(int) error) error {
	raw, err := conn.SyscallConn()
	if err != nil {
		return err
	}
	var fdErr error
	if err := raw.Control(func(fd uintptr) {
		fdErr = cb(int(fd))
	}); err != nil {
		return err
	}
	return fdErr
}
