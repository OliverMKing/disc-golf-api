package log

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

var logger *zap.Logger

const operationIdKey = "operationId"

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("failed to initialize logger: %v", err))
	}
}

// NewContext returns logger from context (creating if necessary) and adds fields to the logger
func NewContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	return context.WithValue(ctx, loggerKey, WithContext(ctx).With(fields...))
}

// NewContextWithOperationId returns logger from context including new operationId field
func NewContextWithOperationId(ctx context.Context, fields ...zapcore.Field) context.Context {
	operationId := uuid.New()
	operationIdField := zap.String(operationIdKey, operationId.String())
	return NewContext(ctx, append(fields, operationIdField)...)
}

// WithContext retrieves and returns the logger from the context
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}

	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}

	return logger
}
