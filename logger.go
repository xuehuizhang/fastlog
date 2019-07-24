package log

import "fmt"

var logger LogInterface

/*
	name
	  file: NewFileLogger
	  console: NewConsoleLogger
*/
func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		logger, err = NewFileLogger(config)
	case "console":
		logger, err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("no support error type %s", name)
	}
	return
}

//Debug 0
func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}

//Trace 1
func Trace(format string, args ...interface{}) {
	logger.Trace(format, args...)
}

//Info 2
func Info(format string, args ...interface{}) {
	logger.Info(format, args...)
}

//Warn 3
func Warn(format string, args ...interface{}) {
	logger.Warn(format, args...)
}

//Error 4
func Error(format string, args ...interface{}) {
	logger.Error(format, args...)
}

//Fatal 5
func Fatal(format string, args ...interface{}) {
	logger.Fatal(format, args...)
}
