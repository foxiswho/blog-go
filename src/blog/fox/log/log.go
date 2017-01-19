package log

import (
	"github.com/astaxie/beego/logs"
	"strings"
)
//此处 为以后 更换框架做准备
// Alert logs a message at alert level.
func Alert(v ...interface{}) {
	logs.Alert(generateFmtStr(len(v)), v...)
}

// Critical logs a message at critical level.
func Critical(v ...interface{}) {
	logs.Critical(generateFmtStr(len(v)), v...)
}

// Error logs a message at error level.
func Error(v ...interface{}) {
	logs.Error(generateFmtStr(len(v)), v...)
}

// Warning logs a message at warning level.
func Warning(v ...interface{}) {
	logs.Warning(generateFmtStr(len(v)), v...)
}

// Warn compatibility alias for Warning()
func Warn(v ...interface{}) {
	logs.Warn(generateFmtStr(len(v)), v...)
}

// Notice logs a message at notice level.
func Notice(v ...interface{}) {
	logs.Notice(generateFmtStr(len(v)), v...)
}

// Informational logs a message at info level.
func Informational(v ...interface{}) {
	logs.Informational(generateFmtStr(len(v)), v...)
}

// Info compatibility alias for Warning()
func Info(v ...interface{}) {
	logs.Info(generateFmtStr(len(v)), v...)
}
// Debug logs a message at debug level.
func Debug(v ...interface{}) {
	logs.Debug(generateFmtStr(len(v)), v...)
}

// Trace logs a message at trace level.
// compatibility alias for Warning()
func Trace(v ...interface{}) {
	logs.Trace(generateFmtStr(len(v)), v...)
}
func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}