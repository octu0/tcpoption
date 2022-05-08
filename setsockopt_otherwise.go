//go:build ignore
// +build ignore

package tcpoption

import (
	"time"
)

func setsockoptLingerTimeout(fd int, d time.Duration) error {
	return nil // not support
}

func setsockoptKeepAliveIdle(fd int, sec int) error {
	return nil // not support
}

func getsockoptKeepAliveIdle(fd int) (int, error) {
	return 0, nil // not support
}

func setsockoptKeepAliveInterval(fd int, sec int) error {
	return nil // not support
}

func getsockoptKeepAliveInterval(fd int) (int, error) {
	return 0, nil // not support
}

func setsockoptKeepAliveProbes(fd int, count int) error {
	return nil // not support
}

func getsockoptKeepAliveProbes(fd int) (int, error) {
	return 0, nil // not support
}

func setsockoptFastOpen(fd int, count int) error {
	return nil // not support
}

func getsockoptFastOpen(fd int) (int, error) {
	return 0, nil // not support
}

func setsockoptFastOpenConnect(fd int, count int) error {
	return nil // not support
}

func getsockoptFastOpenConnect(fd int) (int, error) {
	return 0, nil // not support
}

func setsockoptQuickACK(fd int, onoff int) error {
	return nil // not support
}

func setsockoptDeferAccept(fd int, onoff int) error {
	return nil // not support
}

func setsockoptReuseAddr(fd int, onoff int) error {
	return nil // not support
}

func getsockoptReuseAddr(fd int) (int, error) {
	return 0, nil // not support
}

func setsockoptReusePort(fd int, onoff int) error {
	return nil // not support
}

func getsockoptReusePort(fd int) (int, error) {
	return 0, nil // not support
}
