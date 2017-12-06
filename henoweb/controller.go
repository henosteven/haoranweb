package henoweb

import (
    "net/http"
    "fmt"
)

type Controller interface {
    ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HENOController struct {
   ControllerName string
   Action []string
}

func (c HENOController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello~my~controller")
}
