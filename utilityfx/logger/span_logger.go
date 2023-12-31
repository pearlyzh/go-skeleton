package logger

import (
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type spanLogger struct {
	logger *zap.Logger
	span   trace.Span
}

func (l *spanLogger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, l.getSpanFields(fields)...)
}

func (l *spanLogger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, l.getSpanFields(fields)...)
}

func (l *spanLogger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, l.getSpanFields(fields)...)
}

func (l *spanLogger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, l.getSpanFields(fields)...)
}

func (l *spanLogger) Fatal(msg string, fields ...zapcore.Field) {
	l.logger.Fatal(msg, l.getSpanFields(fields)...)
}

func (l *spanLogger) Panic(msg string, fields ...zapcore.Field) {
	l.logger.Panic(msg, l.getSpanFields(fields)...)
}

func (l *spanLogger) getSpanFields(fields []zapcore.Field) []zapcore.Field {
	traceID := strings.TrimPrefix(l.span.SpanContext().TraceID().String(), "0000000000000000")
	return append(fields,
		zap.String("trace_id", traceID),
		zap.String("span_id", l.span.SpanContext().SpanID().String()),
	)
}
