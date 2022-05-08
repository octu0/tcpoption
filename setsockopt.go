package tcpoption

import (
	"net"
)

func setsockoptLinger(conn *net.TCPConn, sec int) error {
	return conn.SetLinger(sec)
}

func setsockoptReadBuffer(conn *net.TCPConn, bytes int) error {
	return conn.SetReadBuffer(bytes)
}

func setsockoptWriteBuffer(conn *net.TCPConn, bytes int) error {
	return conn.SetWriteBuffer(bytes)
}

func setsockoptNoDelay(conn *net.TCPConn, enable bool) error {
	return conn.SetNoDelay(enable)
}

func setsockoptKeepAlive(conn *net.TCPConn, enable bool) error {
	return conn.SetKeepAlive(enable)
}
