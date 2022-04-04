package logger

import "go.uber.org/zap"

var logger *zap.Logger

func init() {
	SetDebugMode(true)
}

func SetDebugMode(isDebug bool) {
	if isDebug {
		config := zap.NewDevelopmentConfig()
		logger, _ = config.Build(zap.AddCallerSkip(1))
	} else {
		// 構造化ログとして出力される
		config := zap.NewProductionConfig()
		logger, _ = config.Build(zap.AddCallerSkip(1))
	}
	defer logger.Sync()
}

// Debug Wrapper Function
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info Wrapper Function
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn Wrapper Function
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Fatal Wrapper Function
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
