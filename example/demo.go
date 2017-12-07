package main

import (
    "henoweb"
    "syscall"
    "fmt"
    "net/http"
    "os"
)

func main() {
    /* 注册信号处理函数 */
    henoweb.RegisterSignalHook(syscall.SIGINT, recvSigINT)
    henoweb.RegisterSignalHook(syscall.SIGTERM, recvSigTERM)
    
    /* 注册路由 */
    henoweb.RegisterFuncRouter("/example", example)

    indexController := IndexController{}
    henoweb.RegisterControllerRouter("/index", indexController)
    henoweb.Run()
}

func recvSigINT() error {
    fmt.Println("recv sigint ~hi~")
    os.Exit(1)
    return nil
}

func recvSigTERM() error{
    fmt.Println("recv sigterm ~hi~")
    os.Exit(1)
    return nil
}

func example(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello~example~")
}

type IndexController struct {
    henoweb.HENOController
}

func (_ IndexController) Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello~indexc~indexa")
}

func (_ IndexController) Say(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello~indexc~saya")
}
