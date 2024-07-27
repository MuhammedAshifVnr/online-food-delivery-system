package logger

import (
    "log"
    "os"
)

var (
    infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func InitLogger() {
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func LogInfo(v ...interface{}) {
    infoLogger.Println(v...)
}

func LogError(v ...interface{}) {
    errorLogger.Println(v...)
}