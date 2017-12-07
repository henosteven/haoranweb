package henoweb

import (
    "reflect"
    "net/http"
    "fmt"
)

type Controller interface {
}

type HENOController struct {
   ControllerName string
   cp Controller
}

func (_ HENOController) Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome~page")
}

func Invoke(h Controller, methodName string, input ...interface{}) {
    cv := reflect.ValueOf(h)
    var argv = make([]reflect.Value, 0)
    for _, arg := range input {
        argv = append(argv, reflect.ValueOf(arg)) 
    }
    cv.MethodByName(methodName).Call(argv)
}
