# `tcpoption`

[![MIT License](https://img.shields.io/github/license/octu0/tcpoption)](https://github.com/octu0/tcpoption/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/octu0/tcpoption?status.svg)](https://godoc.org/github.com/octu0/tcpoption)
[![Go Report Card](https://goreportcard.com/badge/github.com/octu0/tcpoption)](https://goreportcard.com/report/github.com/octu0/tcpoption)
[![Releases](https://img.shields.io/github/v/release/octu0/tcpoption)](https://github.com/octu0/tcpoption/releases)

additional syscall for tcp socketopt

- `TCP_KEEPIDLE`  / `tcp_keepalive_time`   SetKeepAliveTime
- `TCP_KEEPINTVL` / `tcp_keepalive_intvl`  SetKeepAliveInterval
- `TCP_KEEPCNT`   / `tcp_keepalive_probes` SetKeepAliveProbes
- `TCP_LINGER2`   / `tcp_fin_timeout`      SetLingerTimeout
- `TCP_NODELAY`                            SetNoDelay
- `TCP_FASTOPEN`  / `tcp_fastopen`         SetFastOpen/SetFastOpenConnect
- `TCP_QUICKACK`                           SetQuickACK
- `TCP_DEFER_ACCEPT`                       SetDeferAccept
- `SO_RCVBUF`    SetReadBuffer
- `SO_SNDBUF`    SetWriteBuffer
- `SO_KEEPALIVE` SetKeepAlive
- `SO_LINGER`    SetLinger
- `SO_REUSEADDR` SetReuseAddr
- `SO_REUSEPORT` SetReusePort
