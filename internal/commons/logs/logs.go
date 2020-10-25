package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log   *zap.Logger
	sugar *zap.SugaredLogger
)

// Logger interface to use in every mock
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

/*
InitializeZapCustomLogger Function initializes a logger using uber-go/zap package in the application.
*/
func InitDefault() {
	conf := zap.Config{
		Encoding: "json",
		Level:    zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "file",
			MessageKey:   "msg",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		DisableCaller:    false,
	}
	log, _ = conf.Build()
	sugar = log.Sugar()
}

func Log() Logger {
	_ = log.Sync()
	return log
}

func Sugar() *zap.SugaredLogger {
	_ = sugar.Sync()
	return sugar
}
