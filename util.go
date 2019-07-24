package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

//GetLevelTest 获取日志级别 文字描述
func GetLevelTest(level int) string {
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "Trace"
	case 2:
		return "INFO"
	case 3:
		return "WARN"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	}
	return "NIL"
}

//GetLevelNum 获取日志级别 数字描述
func GetLevelNum(level string) int {
	switch level {
	case "DEBUG":
		return 0
	case "Trace":
		return 1
	case "INFO":
		return 2
	case "WARN":
		return 3
	case "ERROR":
		return 4
	case "FATAL":
		return 5
	}
	return -1
}

//GetLineInfo 获取行号
func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

func writeLog(file *os.File, level int, format string, args ...interface{}) {
	nowStr := time.Now().Format("2006/01/02 15:04:05:999")
	levelTest := GetLevelTest(level)
	fileName, funcName, lineNo := GetLineInfo()

	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)

	fmt.Fprintf(file, "%s %s (%s:%s %d) %s\n", nowStr, levelTest, fileName, funcName, lineNo, msg)
}
