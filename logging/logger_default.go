package logging

import (
	"fmt"
	"log"
)

type DefaultLogger struct {
	minLevel     LogLevel
	loggers      map[LogLevel]*log.Logger
	triggerPanic bool
}

func (l *DefaultLogger) MinLogLevel() LogLevel {
	return l.minLevel
}

func (l *DefaultLogger) write(level LogLevel, message string) {
	if l.minLevel <= level {
		_ = l.loggers[level].Output(2, message)
	}
}

func (l *DefaultLogger) Trace(msg string) {
	l.write(Trace, msg)
}

func (l *DefaultLogger) TraceF(msg string, args ...interface{}) {
	l.write(Trace, fmt.Sprintf(msg, args...))
}

func (l *DefaultLogger) Debug(msg string) {
	l.write(Debug, msg)
}

func (l *DefaultLogger) DebugF(msg string, args ...interface{}) {
	l.write(Debug, fmt.Sprintf(msg, args...))
}

func (l *DefaultLogger) Info(msg string) {
	l.write(Information, msg)
}

func (l *DefaultLogger) InfoF(msg string, args ...interface{}) {
	l.write(Information, fmt.Sprintf(msg, args...))
}

func (l *DefaultLogger) Warn(msg string) {
	l.write(Warning, msg)
}

func (l *DefaultLogger) WarnF(msg string, args ...interface{}) {
	l.write(Warning, fmt.Sprintf(msg, args...))
}

func (l *DefaultLogger) Panic(msg string) {
	l.write(Fatal, msg)
	if l.triggerPanic {
		panic(msg)
	}
}

func (l *DefaultLogger) PanicF(msg string, args ...interface{}) {
	formattedMessage := fmt.Sprintf(msg, args...)
	l.write(Fatal, formattedMessage)
	if l.triggerPanic {
		panic(formattedMessage)
	}
}
