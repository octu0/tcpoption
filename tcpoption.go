package tcpoption

import (
	"net"
	"time"
)

func SetNoLinger(conn net.Conn, enable bool) error {
	if enable {
		if c, ok := conn.(*net.TCPConn); ok {
			return setsockoptLinger(c, 0)
		}
	}
	return nil
}

func SetLinger(conn net.Conn, d time.Duration) error {
	if d == 0 {
		return SetNoLinger(conn, true)
	}
	if c, ok := conn.(*net.TCPConn); ok {
		return setsockoptLinger(c, IntSecond(d))
	}
	return nil
}

func SetLingerTimeout(conn net.Conn, d time.Duration) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return setsockoptLingerTimeout(fd, d)
		})
	}
	return nil
}

func SetReadBuffer(conn net.Conn, bytes int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return setsockoptReadBuffer(c, bytes)
	}
	return nil
}

func SetWriteBuffer(conn net.Conn, bytes int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return setsockoptWriteBuffer(c, bytes)
	}
	return nil
}

func SetNoDelay(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return setsockoptNoDelay(c, enable)
	}
	return nil
}

func KeepAlive(conn net.Conn, enable bool, idle, interval time.Duration, probes int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := setsockoptKeepAlive(c, enable); err != nil {
			return err
		}
		if enable != true {
			return nil
		}
		return getFd(c, func(fd int) error {
			if err := setsockoptKeepAliveIdle(fd, IntSecond(idle)); err != nil {
				return err
			}
			if err := setsockoptKeepAliveInterval(fd, IntSecond(interval)); err != nil {
				return err
			}
			return setsockoptKeepAliveProbes(fd, probes)
		})
	}
	return nil
}

func SetKeepAlive(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return setsockoptKeepAlive(c, enable)
	}
	return nil
}

func SetKeepAliveTime(conn net.Conn, d time.Duration) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := setsockoptKeepAlive(c, true); err != nil {
			return err
		}
		return getFd(c, func(fd int) error {
			return setsockoptKeepAliveIdle(fd, IntSecond(d))
		})
	}
	return nil
}

func SetKeepAliveInterval(conn net.Conn, d time.Duration) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := setsockoptKeepAlive(c, true); err != nil {
			return err
		}
		return getFd(c, func(fd int) error {
			return setsockoptKeepAliveInterval(fd, IntSecond(d))
		})
	}
	return nil
}

func SetKeepAliveProbes(conn net.Conn, count int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		if err := setsockoptKeepAlive(c, true); err != nil {
			return err
		}
		return getFd(c, func(fd int) error {
			return setsockoptKeepAliveProbes(fd, count)
		})
	}
	return nil
}

func SetFastOpen(conn net.Conn, count int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return setsockoptFastOpen(fd, count)
		})
	}
	return nil
}

func SetFastOpenFd(fd int, count int) error {
	return setsockoptFastOpenConnect(fd, count)
}

func SetFastOpenConnect(conn net.Conn, count int) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return setsockoptFastOpenConnect(fd, count)
		})
	}
	return nil
}

func SetFastOpenConnectFd(fd int, count int) error {
	return setsockoptFastOpenConnect(fd, count)
}

func SetQuickACK(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return setsockoptQuickACK(fd, IntBool(enable))
		})
	}
	return nil
}

func SetQuickACKFd(fd int, enable bool) error {
	return setsockoptQuickACK(fd, IntBool(enable))
}

func SetDeferAccept(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return setsockoptDeferAccept(fd, IntBool(enable))
		})
	}
	return nil
}

func SetDeferAcceptFd(fd int, enable bool) error {
	return setsockoptDeferAccept(fd, IntBool(enable))
}

func SetReuseAddr(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return setsockoptReuseAddr(fd, IntBool(enable))
		})
	}
	return nil
}

func SetReuseAddrFd(fd int, enable bool) error {
	return setsockoptReuseAddr(fd, IntBool(enable))
}

func SetReusePort(conn net.Conn, enable bool) error {
	if c, ok := conn.(*net.TCPConn); ok {
		return getFd(c, func(fd int) error {
			return setsockoptReusePort(fd, IntBool(enable))
		})
	}
	return nil
}

func SetReusePortFd(fd int, enable bool) error {
	return setsockoptReusePort(fd, IntBool(enable))
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
			if err := setsockoptLinger(c, 0); err != nil {
				return err
			}
		}
		if 0 < cfg.LingerTimeout {
			if err := setsockoptLingerTimeout(fd, cfg.LingerTimeout); err != nil {
				return err
			}
		}
		if 0 < cfg.ReadBuffer {
			if err := setsockoptReadBuffer(c, cfg.ReadBuffer); err != nil {
				return err
			}
		}
		if 0 < cfg.WriteBuffer {
			if err := setsockoptWriteBuffer(c, cfg.WriteBuffer); err != nil {
				return err
			}
		}
		if err := setsockoptNoDelay(c, cfg.EnableNoDelay); err != nil {
			return err
		}
		if err := setsockoptKeepAlive(c, cfg.EnableKeepAlive); err != nil {
			return err
		}
		if cfg.EnableKeepAlive {
			if 0 < cfg.KeepAliveTime {
				if err := setsockoptKeepAliveIdle(fd, IntSecond(cfg.KeepAliveTime)); err != nil {
					return err
				}
			}
			if 0 < cfg.KeepAliveInterval {
				if err := setsockoptKeepAliveInterval(fd, IntSecond(cfg.KeepAliveInterval)); err != nil {
					return err
				}
			}
			if 0 < cfg.KeepAliveProbes {
				if err := setsockoptKeepAliveProbes(fd, cfg.KeepAliveProbes); err != nil {
					return err
				}
			}
		}
		if 0 < cfg.FastOpen {
			if err := setsockoptFastOpen(fd, cfg.FastOpen); err != nil {
				return err
			}
		}
		if 0 < cfg.FastOpenConnect {
			if err := setsockoptFastOpenConnect(fd, cfg.FastOpenConnect); err != nil {
				return err
			}
		}
		if err := setsockoptQuickACK(fd, IntBool(cfg.EnableQuickACK)); err != nil {
			return err
		}
		if err := setsockoptDeferAccept(fd, IntBool(cfg.EnableDeferAccept)); err != nil {
			return err
		}
		if err := setsockoptReuseAddr(fd, IntBool(cfg.EnableReuseAddr)); err != nil {
			return err
		}
		if err := setsockoptReusePort(fd, IntBool(cfg.EnableReusePort)); err != nil {
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
