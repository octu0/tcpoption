package tcpoption

import (
	"net"
)

func tcpSetLinger(conn *net.TCPConn, sec int) error {
	return conn.SetLinger(sec)
}

func tcpSetReadBuffer(conn *net.TCPConn, bytes int) error {
	return conn.SetReadBuffer(bytes)
}

func tcpSetWriteBuffer(conn *net.TCPConn, bytes int) error {
	return conn.SetWriteBuffer(bytes)
}

func tcpSetNoDelay(conn *net.TCPConn, enable bool) error {
	return conn.SetNoDelay(enable)
}

func tcpSetKeepAlive(conn *net.TCPConn, enable bool) error {
	return conn.SetKeepAlive(enable)
}

/* Implemented for each OS
func tcpSetLingerTimeout(fd int, d time.Duration) error
*/

/* Implemented for each OS
func tcpSetKeepAliveIdle(fd int, sec int) error
*/

/* Implemented for each OS
func tcpSetKeepAliveInterval(fd int, sec int) error
*/

/* Implemented for each OS
func tcpSetKeepAliveProbes(fd int, count int) error
*/

/* Implemented for each OS
func tcpSetFastOpen(fd int, count int) error
*/

/* Implemented for each OS
func tcpSetFastOpenConnect(fd int, count int) error
*/

/* Implemented for each OS
func tcpSetQuickACK(fd int, onoff int) error
*/

/* Implemented for each OS
func tcpSetDeferAccept(fd int, onoff int) error
*/

/* Implemented for each OS
func tcpSetReuseAddr(fd int, onoff int) error
*/

/* Implemented for each OS
func tcpSetReusePort(fd int, onoff int) error
*/
