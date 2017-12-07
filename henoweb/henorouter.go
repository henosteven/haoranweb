package henoweb

import (
    "net/http"
    "fmt"
)

type HenoRouter struct {
   HandleMap map[string]interface{} 
}

var hRouter = NewRouter()

func NewRouter() *HenoRouter{
    r := &HenoRouter{}
    r.HandleMap = make(map[string]interface{})
    return r 
}

func RegisterFuncRouter(pattern string, h http.HandlerFunc) {
    hRouter.HandleMap[pattern] = h
}

func RegisterControllerRouter(pattern string, c Controller) {
    hRouter.HandleMap[pattern] = c
}

func (router *HenoRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h, ok := router.HandleMap[r.URL.Path]
    if !ok {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "~opps 404~")
    } else {
        switch h.(type) {
            case http.HandlerFunc:
                h.(http.HandlerFunc).ServeHTTP(w, r) 
            case Controller:
                fmt.Println(r.URL)
                Invoke(h, "Index", w, r)
        }
    }
}
