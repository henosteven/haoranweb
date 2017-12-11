package cache

import (
    "github.com/garyburd/redigo/redis"
    "fmt"
)

type CacheConfig struct {
    Addr string 
}

var RedisConn redis.Conn 

func InitRedisServer(addr string) {
    var err error
    RedisConn, err = redis.Dial("tcp", addr)
    if err != nil {
        fmt.Println(err)
    }
}

func Set(k, v string) {
    RedisConn.Do("SET", k, v)
}

func Get(k string) string{
    v, err := redis.String(RedisConn.Do("GET", k))
    if err != nil {
        fmt.Println(err)
    }
    return v
}

func Hset(k, field, v string) {
    RedisConn.Do("HSET", k, field, v)
}

func Hget(k, field string) string{
    v, err := RedisConn.Do("HGET", k, field)
    if err != nil {
        fmt.Println(err)
    }
    result, _ := v.([]byte)
    return string(result)
}

func Hgetall(k string) map[string]string{
    v, err := RedisConn.Do("HGETALL", k)
    if err != nil {
        fmt.Println(err)
    }
    result, _ := v.([]interface{})
    var stringlist  = make(map[string]string, 0)
    if len(result) > 0 {
        var curKey string = ""
        for _, tmp := range result {
            tmpp, _ := tmp.([]byte)
            if curKey == "" {
                curKey = string(tmpp)
            } else {
                stringlist[curKey] = string(tmpp)
                curKey = ""
            }
        }
    } 
    return stringlist
}
