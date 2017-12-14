package henoweb

import (
	"net"
)

type HenoListener struct {
	net.Listener
}

func NewHenoListener(ln net.Listener) *HenoListener {
	var hl = &HenoListener{Listener: ln}
	return hl
}

func (ln *HenoListener) Accept() (conn net.Conn, err error) {
	conn, err = ln.Listener.Accept()
	if err != nil {
		return
	}
	conn = HenoConn{Conn: conn}
	return
}
