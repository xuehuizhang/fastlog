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
	logdata := writeLog(LogLevelDebug, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s %d) %s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

//Trace 1
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	logdata := writeLog(LogLevelTrace, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s %d) %s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

//Info 2
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	logdata := writeLog(LogLevelInfo, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s %d) %s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

//Warn 3
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	logdata := writeLog(LogLevelWarn, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s %d) %s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

//Error 4
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	logdata := writeLog(LogLevelError, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s %d) %s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

//Fatal 5
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	logdata := writeLog(LogLevelFatal, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s (%s:%s %d) %s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

//Close 释放资源
func (c *ConsoleLogger) Close() {

}
