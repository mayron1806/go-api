package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug  *log.Logger
	info   *log.Logger
	warn   *log.Logger
	err    *log.Logger
	writer io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)

	return &Logger{
		writer: writer,
		debug:  log.New(writer, prefix+"[DEBUG] ", log.Ltime|log.Ldate),
		info:   log.New(writer, prefix+"[INFO] ", log.Ltime|log.Ldate),
		warn:   log.New(writer, prefix+"[WARN] ", log.Ltime|log.Ldate),
		err:    log.New(writer, prefix+"[ERROR] ", log.Ltime|log.Ldate),
	}
}
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.warn.Println(v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warn.Printf(format, v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
