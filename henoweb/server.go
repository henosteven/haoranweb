/*
 Package henoweb provides a webserver framework
 import package henoweb, then just run~~
 everything will be ok
*/
package henoweb

import (
	"net"
	"net/http"
)

// Run start moniter the signal
// register the router
// then start to serve
func Run() {
	go RunSignalHook()
	server := http.Server{Addr: ":8080", Handler: hRouter}
	ln, _ := net.Listen("tcp", server.Addr)
	hl := NewHenoListener(ln)
	server.Serve(hl)
}
