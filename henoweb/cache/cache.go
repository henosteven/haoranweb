package cache

import (
    "github.com/garyburd/redigo/redis"
    "fmt"
)

type CacheConfig struct {
    Addr string 
}

var RedisConn redis.Conn 

func InitRedisServer() {
    var err error
    RedisConn, err = redis.Dial("tcp", "10.94.107.14:6379")
    if err != nil {
        fmt.Println(err)
    }
}

func Set(k, v string) {
    RedisConn.Do("SET", k, v)
}

func Get(k string) string{
    var err error
    v, err := redis.String(RedisConn.Do("GET", k))
    if err != nil {
        fmt.Println(err)
    }
    return v
}
