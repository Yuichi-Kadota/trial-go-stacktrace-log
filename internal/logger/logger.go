package logger

import (
	"time"

	pe "github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

// stacktrace付きでエラーを出力する
func Error(err error) {
	// pkg/errorsのstacktraceを持ったerrにUnwrapする
	log.Error().Stack().Err(pe.Unwrap(err)).Msg(err.Error())
}
