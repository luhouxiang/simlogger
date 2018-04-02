package simlogger

import (
	"fmt"
	"log"
	"os"
)

type LEVEL int32

var logLevel LEVEL = 1
var logFileName string = ""

const (
	ALL LEVEL = iota
	TRACE
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

var console_trace_log *log.Logger
var console_debug_log *log.Logger
var console_info_log *log.Logger
var console_warn_log *log.Logger
var console_error_log *log.Logger
var console_fatal_log *log.Logger

var file_trace_log *log.Logger
var file_debug_log *log.Logger
var file_info_log *log.Logger
var file_warn_log *log.Logger
var file_error_log *log.Logger
var file_fatal_log *log.Logger

func SetLevel(level int) {
	if level < int(ALL) {
		logLevel = ALL
	}
	if level > int(FATAL) {
		logLevel = OFF
	}
	logLevel = LEVEL(level)
}

func SetFileName(filename string) {
	logFile, err := os.Create(filename)
	if err != nil {
		return
	}
	file_trace_log = log.New(logFile, "[trace]", log.Lshortfile)
	file_debug_log = log.New(logFile, "[debug]", log.Lshortfile)
	file_info_log = log.New(logFile, "[info]", log.Lshortfile)
	file_warn_log = log.New(logFile, "[warn]", log.Lshortfile)
	file_error_log = log.New(logFile, "[error]", log.Lshortfile)
	file_fatal_log = log.New(logFile, "[fatal]", log.Lshortfile)

}

func init() {
	logLevel = ALL
	console_trace_log = log.New(os.Stdout, "[trace]", log.Lshortfile)
	console_debug_log = log.New(os.Stdout, "[debug]", log.Lshortfile)
	console_info_log = log.New(os.Stdout, "[info]", log.Lshortfile)
	console_warn_log = log.New(os.Stdout, "[warn]", log.Lshortfile)
	console_error_log = log.New(os.Stdout, "[error]", log.Lshortfile)
	console_fatal_log = log.New(os.Stdout, "[fatal]", log.Llongfile)
}

func Trace(format string, args ...interface{}) {
	if logLevel <= TRACE {
		console_trace_log.Output(2, fmt.Sprintf(format, args...))
		file_trace_log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Debug(format string, args ...interface{}) {
	if logLevel <= DEBUG {
		console_debug_log.Output(2, fmt.Sprintf(format, args...))
		if file_debug_log != nil {
			file_debug_log.Output(2, fmt.Sprintf(format, args...))
		}
	}
}

func Info(format string, args ...interface{}) {
	if logLevel <= INFO {
		console_info_log.Output(2, fmt.Sprintf(format, args...))
		if file_info_log != nil {
			file_info_log.Output(2, fmt.Sprintf(format, args...))
		}
	}
}

func Warn(format string, args ...interface{}) {
	if logLevel <= WARN {
		console_warn_log.Output(2, fmt.Sprintf(format, args...))
		if file_warn_log != nil {
			file_warn_log.Output(2, fmt.Sprintf(format, args...))
		}
	}
}

func Error(format string, args ...interface{}) {
	if logLevel <= ERROR {
		console_error_log.Output(2, fmt.Sprintf(format, args...))
		if file_error_log != nil {
			file_error_log.Output(2, fmt.Sprintf(format, args...))
		}
	}
}

func Fatal(format string, args ...interface{}) {
	if logLevel <= FATAL {
		console_fatal_log.Output(2, fmt.Sprintf(format, args...))
		if file_fatal_log != nil {
			file_fatal_log.Output(2, fmt.Sprintf(format, args...))
		}
	}
}

func main() {
	Info("warn abcdef")
	Warn("warn abcdef")
}
