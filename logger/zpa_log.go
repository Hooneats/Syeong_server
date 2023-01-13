package logger

import (
	"bytes"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
)

var instance *zapLog

type LogConfigure interface {
	GetSettingValues() (path, level string, size, backup, age int)
}

type zapLog struct {
	lg *zap.Logger
}

func NewZapLog() *zapLog {
	if instance != nil {
		return instance
	}
	instance = &zapLog{
		lg: zap.L(),
	}
	return instance
}

// registerGlobalZapLogger global zap L 및 SugaredLogger 를 대체
func RegisterGlobalZapLogger(lcfg LogConfigure) {
	path, level, size, backup, age := lcfg.GetSettingValues()

	now := time.Now()
	lPath := fmt.Sprintf("%s_%s.log", path, now.Format("2006-01-02"))
	// 설정 옵션
	writeSyncer := getLogWriter(lPath, size, backup, age)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(level)); err != nil {
		panic("logger load, fail")
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)
	z := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(z)
}

func (z *zapLog) Debug(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	z.lg.Debug("debug", zap.String("-", b.String()))
}

// Info is a convenient alias for Root().Info
func (z *zapLog) Info(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	z.lg.Info("info", zap.String("-", b.String()))
}

// Warn is a convenient alias for Root().Warn
func (z *zapLog) Warn(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	z.lg.Warn("warn", zap.String("-", b.String()))
}

// Error is a convenient alias for Root().Error
func (z *zapLog) Error(ctx ...interface{}) {
	b := z.newWrittenBuffer(ctx)

	z.lg.Error("error", zap.String("-", b.String()))
}

func (z *zapLog) newWrittenBuffer(ctx []interface{}) *bytes.Buffer {
	var b bytes.Buffer
	for _, str := range ctx {
		b.WriteString(str.(string))
		b.WriteString(" ")
	}
	return &b
}

// encoder 옵션 설정
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
