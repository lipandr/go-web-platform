package logging

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

type Logger interface {
	Trace(string)
	TraceF(string, ...interface{})
	Debug(string)
	DebugF(string, ...interface{})
	Info(string)
	InfoF(string, ...interface{})
	Warn(string)
	WarnF(string, ...interface{})
	Panic(string)
	PanicF(string, ...interface{})
}
