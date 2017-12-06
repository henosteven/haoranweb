package henoweb

import (
    "net"
    "net/http"
)

func Run() {
   go RunSignalHook()
   server := http.Server{Addr:":8080", Handler:hRouter}
   ln, _ := net.Listen("tcp", server.Addr)
   hl := NewHenoListener(ln)
   server.Serve(hl)
}
