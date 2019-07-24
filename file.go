package log

import (
	"fmt"
	"os"
	"strconv"
)

//FileLogger 文件 Logger
type FileLogger struct {
	level       int
	logPath     string
	logName     string
	file        *os.File
	warnFile    *os.File
	logDataChan chan *LogData
}

//NewFileLogger 实例化一个File Logger
func NewFileLogger(config map[string]string) (logger LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log level")
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log name")
		return
	}

	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log path")
		return
	}
	logLevelNum := GetLevelNum(logLevel)

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}

	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}

	logger = &FileLogger{
		level:       logLevelNum,
		logPath:     logPath,
		logName:     logName,
		logDataChan: make(chan *LogData, chanSize),
	}
	logger.Init()
	return
}

//Init 初始化函数
func (f *FileLogger) Init() {
	//初始化debug trace info 文件句柄
	fileName := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed,err:%v", fileName, err))
	}
	f.file = file

	//初始化warn error fatal 文件句柄
	fileName = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed,err:%v", fileName, err))
	}
	f.warnFile = file

	go f.writeLogForChan()
}

func (f *FileLogger) writeLogForChan() {
	for logdata := range f.logDataChan {
		var file = f.file
		if logdata.WarnOrFatal {
			file = f.warnFile
		}
		fmt.Fprintf(file, "%s %s (%s:%s %d) %s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
	}
}

//SetLevel 设置日志级别
func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

//Debug Debug
func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	datalog := writeLog(LogLevelDebug, format, args...)
	select {
	case f.logDataChan <- datalog:
	default:
	}

}

//Trace Trace
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	datalog := writeLog(LogLevelTrace, format, args...)
	select {
	case f.logDataChan <- datalog:
	default:
	}
}

//Info Info
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	datalog := writeLog(LogLevelInfo, format, args...)
	select {
	case f.logDataChan <- datalog:
	default:
	}
}

//Warn Warn
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	datalog := writeLog(LogLevelWarn, format, args...)
	select {
	case f.logDataChan <- datalog:
	default:
	}
}

//Error Error
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	datalog := writeLog(LogLevelError, format, args...)
	select {
	case f.logDataChan <- datalog:
	default:
	}
}

//Fatal Fatal
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	datalog := writeLog(LogLevelFatal, format, args...)
	select {
	case f.logDataChan <- datalog:
	default:
	}
}

//Close 关闭文件句柄
func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
