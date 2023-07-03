package logger

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type ILogger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Panic(msg string, fields ...zapcore.Field)
}

var globalLogger *zapLogger

func init() {
	//default logger.go
	logger := initZapLogger(false)
	globalLogger = &zapLogger{
		logger: logger,
	}
}

func InitOverriddenLogger(production bool) {
	if production {
		logger := initZapLogger(true)
		globalLogger.logger = logger
	}
}

func initZapLogger(production bool) *zap.Logger {
	config := initConfig(production)
	// AddCallerSkip to skip report wrapper as caller in log message
	zapLogger, err := config.Build(
		zap.AddCallerSkip(1),
	)
	if err != nil {
		log.Fatal("Can not create zap-logger!", err)
	}
	return zapLogger
}

func L() ILogger {
	return globalLogger
}

func Ctx(ctx context.Context) ILogger {
	span := trace.SpanFromContext(ctx)
	return &spanLogger{
		logger: globalLogger.logger,
		span:   span,
	}
}
