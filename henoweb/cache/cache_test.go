package cache

import (
    "testing"
)

func TestGet(t *testing.T) {
    InitRedisServer()
    Set("user", "jinjing")
    user := Get("user")
    if user != "jinjing" {
        t.Errorf("expect: %s  , real: %s", "jinjing", user)
    }
}
