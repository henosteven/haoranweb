package logger

import (
    "log"
    "time"
    "io"
    "fmt"
    "os"
)

const (
    _ = iota
    Day
    Hour
    Minute
)

type LogConfig struct {
    Duration int32
    Unit int8
    CurrentPath string
    CurrentOutput io.Writer 
    CurrentLogFile string
}

var CurrentLogConfig LogConfig

func LogInit(path string, duration int32, unit int8) {
    //todo 数据校验
    CurrentLogConfig.Duration = duration
    CurrentLogConfig.Unit = unit
    CurrentLogConfig.CurrentPath = path 
}

func LogTrace(params interface{}) {
    str := "[TRACE]" + formatLogData(params)
    logData(str)
}

func LogNotice(params []interface{}) {
    str := "[NOTICE]" + formatLogData(params)
    logData(str)
}

func LogError(params []interface{}) {
    str := "[ERROR]" + formatLogData(params)
    logData(str)
}

func formatLogData(params interface{}) string {
    return fmt.Sprintf("%v", params)
}

func logData(logstr string) {
    initLogFile()
    log.Print(logstr)
}

func GenTraceID() {
    
}

func initLogFile() {
    var file string
    switch CurrentLogConfig.Unit {
        case Day:
            file = time.Now().Format("20060102")
        case Hour:
            file = time.Now().Format("2006010215")
        case Minute:
            file = time.Now().Format("200601021504")
    }
    
    //todo 需要关闭文件资源
    //需要处理竞争问题
    if file != CurrentLogConfig.CurrentLogFile {
        CurrentLogConfig.CurrentLogFile = file
        fp, _ := os.Create(CurrentLogConfig.CurrentPath + "/" + file)
        CurrentLogConfig.CurrentOutput = fp
        log.SetOutput(fp)
    }
}
