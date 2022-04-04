package logger

import (
	"trial-go-stacktrace/internal/errors"

	"go.uber.org/zap"
)

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

// zapのスタックトレースだと、loggerを呼び出したhandler.goのスタックしか表示されない
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// zapのスタックトレースだとエラー発生箇所のstackが取得できない？ので明示的にstacktraceを追加する
func WarnErr(err error) {
	st := errors.StackTrace(err)
	logger.Warn(err.Error(), zap.String("Stack", st))
}
