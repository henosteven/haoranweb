package henoweb

import (
    "net"
    "net/http"
    "fmt"
)

func Run() {
   go RunSignalHook()
   serveMux := http.NewServeMux()  
   serveMux.HandleFunc("/example", example)
   server := http.Server{Addr:":8080", Handler:serveMux}
   ln, _ := net.Listen("tcp", server.Addr)
   hl := NewHenoListener(ln)
   server.Serve(hl)
}

func example(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello~henoweb~")
}
