package henoweb

import (
    "os"
    "syscall"
    "os/signal"
    "fmt"
)

var SignalHook = make(map[os.Signal][]SignalHookHandler)

type SignalHookHandler func() error

func RegisterSignalHook(sig os.Signal, sigHandler func() error) {
    handleSlice, ok := SignalHook[sig]
    if !ok {
        handleSlice = make([]SignalHookHandler, 0)
    }
    handleSlice = append(handleSlice, SignalHookHandler(sigHandler))
    SignalHook[sig] = handleSlice
}

func RunSignalHook() {
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT)
    signal.Notify(ch, syscall.SIGTERM)
    signal.Notify(ch, syscall.SIGHUP)

    for {
        sig := <-ch
        fmt.Println(sig)
        handleSlice, ok := SignalHook[sig]
        if !ok || len(handleSlice) == 0 {
            continue
        }

        fmt.Println(handleSlice, SignalHook)
        for _, f := range handleSlice {
            if f != nil {
                f()
            }
        }
    }
}
