package conf

import (
	"context"

	"go.uber.org/zap"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

var logger *zap.Logger
var devLogger *zap.Logger

var pLogger *zap.Logger

// IsDebug デバックログ出力有無
var IsDebug = false

// WithContext WithContext関数
// 役割：
func WithContext(ctx context.Context) *zap.Logger {
	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}
