package logger

import (
	"fmt"

	"github.com/blendle/zapdriver"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	viper.SetDefault("LOG_LEVEL", "debug")
}

var zapLevelMap = map[string]zapcore.Level{
	"debug":    zapcore.DebugLevel,
	"info":     zapcore.InfoLevel,
	"warn":     zapcore.WarnLevel,
	"error":    zapcore.ErrorLevel,
	"panic":    zapcore.PanicLevel,
	"fatal":    zapcore.FatalLevel,
	"devPanic": zapcore.DPanicLevel,
}

var logger *Logger

type Logger struct {
	*zap.Logger
}

func getLevel(level string) zapcore.Level {
	if logLevel, ok := zapLevelMap[level]; ok {
		return logLevel
	}

	panic("invalid log level")
}

func init() {
	productionConfig := zapdriver.NewProductionConfig()
	productionConfig.Level.SetLevel(getLevel(viper.GetString("LOG_LEVEL")))

	zapLogger, err := productionConfig.Build(zapdriver.WrapCore(
		zapdriver.ReportAllErrors(true),
	))

	if err != nil {
		panic(fmt.Errorf("failed to init logger with: %w", err))
	}

	logger = &Logger{
		Logger: zapLogger,
	}
}

func Log() *Logger {
	return logger
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	fmt.Println(msg)
	l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Info(fmt.Sprintf(format, v...))
}
