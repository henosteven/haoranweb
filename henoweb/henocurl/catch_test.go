package henocurl

import (
   "testing" 
   "log"
)

func TestGet(t *testing.T) {
    response := Get("http://www.baidu.com", "10ms")
    log.Print(response)
}
