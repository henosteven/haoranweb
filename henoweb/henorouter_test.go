package henoweb

import (
    "testing"
)

func TestUcfirst(t *testing.T) {
    cases := []struct{
        in, want string
    } {
        {"abc", "Abc"},
        {"Abc", "Abc"},
        {"Zab", "Zab"},
        {"zab", "Zab"},
    }

    for _, item := range cases {
       got := Ucfirst(item.in)
       if got != item.want {
           t.Errorf("in: %s, out: %s, want:%s\n", item.in, got, item.want)
       } 
    }
}

func TestgetControllerAction(t *testing.T) {
    cases := []struct{
        in, action, controller string
    } {
        {"index", "index", "index"},
        {"/index", "index", "index"},
        {"/", "index", "index"},
        {"/heno", "index", "heno"},
        {"/heno/do", "do", "heno"},
        {"/heno/do/xxx", "do", "heno"},
    }

    for _, item := range cases {
       action, controller := getControllerAction(item.in)
       if action != item.action || controller != item.controller {
           t.Errorf("in: %s, action: %s, controller:%s\n", item.in, item.action, item.controller)
       } 
    }
}
