package log

import (
	"fmt"
	"os"
)

//ConsoleLogger 控制台日志
type ConsoleLogger struct {
	level int
}

//NewConsoleLogger 构造函数
func NewConsoleLogger(config map[string]string) (logger LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("no found log level")
		return
	}
	level := GetLevelNum(logLevel)
	logger = &ConsoleLogger{
		level: level,
	}
	return
}

//Init 初始化函数
func (c *ConsoleLogger) Init() {

}

//SetLevel 设置日志级别
func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	c.level = level
}

//Debug 0
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	writeLog(os.Stdout, LogLevelDebug, format, args...)
}

//Trace 1
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	writeLog(os.Stdout, LogLevelTrace, format, args...)
}

//Info 2
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	writeLog(os.Stdout, LogLevelInfo, format, args...)
}

//Warn 3
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	writeLog(os.Stdout, LogLevelWarn, format, args...)
}

//Error 4
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	writeLog(os.Stdout, LogLevelError, format, args...)
}

//Fatal 5
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	writeLog(os.Stdout, LogLevelFatal, format, args...)
}

//Close 释放资源
func (c *ConsoleLogger) Close() {

}
