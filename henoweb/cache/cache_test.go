package cache

import (
    "testing"
)

func TestGet(t *testing.T) {
    InitRedisServer()
    Set("user", "jinjing")
    Hset("jinjing", "name", "jinjing")
    Hset("jinjing", "age", "28")

    age := Hget("jinjing", "age")
    if age != "28" {
        t.Errorf("expect: %s  , real: %s", "28", age)
    }
Hgetall("jinjing")
    user := Get("user")
    if user != "jinjing" {
        t.Errorf("expect: %s  , real: %s", "jinjing", user)
    }
}
