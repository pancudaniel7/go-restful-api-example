package utils

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Logger interface {
	Info(msg string, fields ...any)
	Error(msg string, fields ...any)
	Debug(msg string, fields ...any)
	Warn(msg string, fields ...any)
}

var (
	zapLoggerInstance *ZapLogger
	once              sync.Once
)

type ZapLogger struct {
	mu     sync.Mutex
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// Log Inject what Implementation of logger you will like here that will support the Logger interface
func Log() Logger {
	once.Do(func() {
		zapLoggerInstance = &ZapLogger{}
		zapLoggerInstance.initLogger(viper.GetString("server.log.level"))
	})
	return zapLoggerInstance
}

func (l *ZapLogger) Info(msg string, fields ...any) {
	l.sugar.Infof(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...any) {
	l.sugar.Errorf(msg, fields...)
}

func (l *ZapLogger) Debug(msg string, fields ...any) {
	l.sugar.Debugf(msg, fields...)
}

func (l *ZapLogger) Warn(msg string, fields ...any) {
	l.sugar.Warnf(msg, fields...)
}

func (l *ZapLogger) initLogger(level string) {
	var config zap.Config

	switch {
	case level == "info":
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	case level == "debug":
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	default:
		panic("Invalid log level")
	}

	var err error
	l.logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	l.sugar = l.logger.Sugar()
}
