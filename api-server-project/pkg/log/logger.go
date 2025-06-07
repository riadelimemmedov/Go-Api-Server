package logger

import (
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	err     error
	log     *zap.Logger
	logPath string
	once    sync.Once
	mu      sync.Mutex
	appEnv  = os.Getenv("APP_ENV")
)

// ! Initialize sets up the global logger
func Initialize(env string) *zap.Logger {
	once.Do(func() {
		var config zap.Config = zap.NewProductionConfig()

		if env == "production" {
			logPath = "logs/production"
			config.OutputPaths = []string{filepath.Join(logPath, "app.log")}
		} else {
			logPath = "logs/development"
			config.OutputPaths = []string{"stdout", filepath.Join(logPath, "app.log")}
		}

		if err = os.MkdirAll(logPath, 0755); err != nil {
			panic(err)
		}

		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.StacktraceKey = "stacktrace"

		log, err = config.Build(zap.AddCallerSkip(1))
		if err != nil {
			panic(err)
		}
		zap.ReplaceGlobals(log)
	})
	return log
}

// ! GetLogger returns the global logger instance, initializing it if necessary
func GetLogger() *zap.Logger {
	mu.Lock()
	defer mu.Unlock()

	if log == nil {
		Initialize(appEnv)
	}
	return log
}

// ! Info logs an informational message with optional fields
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// ! Error logs a debug message with optional fields
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// ! Warn logs a warning message
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

// ! Sync buffered logs to disk
func Sync() {
	mu.Lock()
	defer mu.Unlock()

	if log != nil {
		_ = log.Sync()
	}
}
