package logger

type Logger interface {
	Debug(ctx ...interface{})
	Info(ctx ...interface{})
	Warn(ctx ...interface{})
	Error(ctx ...interface{})
}
