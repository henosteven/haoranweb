package henocurl

import (
   "testing" 
   "log"
)

func TestGet(t *testing.T) {
    response := Get("http://www.baidu.com")
    log.Print(response)
}
