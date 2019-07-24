package log

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.SetLevel(0)
	logger.Debug("user id [%d] is come from china", 2322)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()
}
