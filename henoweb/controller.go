package henoweb

import (
    "reflect"
)

type Controller interface {
}

type HENOController struct {
   ControllerName string
   cp Controller
}

func Invoke(h Controller, methodName string, input ...interface{}) {
    cv := reflect.ValueOf(h)
    var argv = make([]reflect.Value, 0)
    for _, arg := range input {
        argv = append(argv, reflect.ValueOf(arg)) 
    }
    cv.MethodByName(methodName).Call(argv)
}
