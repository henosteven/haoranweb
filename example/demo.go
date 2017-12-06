package main

import (
    "henoweb"
    "syscall"
    "fmt"
)

func main() {
    henoweb.RegisterSignalHook(syscall.SIGINT, recvSigINT)
    henoweb.RegisterSignalHook(syscall.SIGTERM, recvSigTERM)
    henoweb.Run()
}

func recvSigINT() error {
    fmt.Println("recv sigint ~hi~")
    return nil
}

func recvSigTERM() error{
    fmt.Println("recv sigterm ~hi~")
    return nil
}
