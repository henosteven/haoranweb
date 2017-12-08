package logger

import (
    "testing"
)

func TestLogInit(t *testing.T) {
   LogInit("/tmp/", 1, Hour) 
   LogTrace("heno")
   LogTrace([]int{1,2,3,4})
}
