package logger

var AppLog Logger

func LoadLogger(conf LogConfigure) {
	RegisterGlobalZapLogger(conf)
	AppLog = NewZapLog()
}
