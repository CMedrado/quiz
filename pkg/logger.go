package pkg

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger initializes and returns a Zap logger instance.
func CreateLogger() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     map[string]interface{}{"pid": os.Getpid()},
	}

	logger, err := config.Build()
	if err != nil {
		panic("failed to create logger")
	}
	return logger
}

type ctxKey struct{}

// WithCtx adds a Zap logger to the context.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok && lp == l {
		return ctx
	}
	return context.WithValue(ctx, ctxKey{}, l)
}

// FromCtx retrieves a Zap logger from the context, or returns a no-op logger if none is found.
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	}
	return zap.NewNop()
}
