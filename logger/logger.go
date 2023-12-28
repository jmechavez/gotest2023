package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	// Configure Zap logger with production settings
	config := zap.NewProductionConfig()

	// Customize encoder configuration
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	// Build the logger
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// Info logs informational messages.
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Debug logs debug messages.
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Error logs error messages.
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}
