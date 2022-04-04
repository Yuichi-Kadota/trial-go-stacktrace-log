package errors

import (
	"fmt"
	"trial-go-stacktrace/codes"

	"github.com/pkg/errors"
)

type privateError struct {
	code codes.Code
	err  error
}

func (e privateError) Error() string {
	return fmt.Sprintf("Code: %s, Msg: %s", e.code, e.err)
}

func (e privateError) Unwrap() error {
	return e.err
}

func Errorf(c codes.Code, format string, a ...interface{}) error {
	if c == codes.OK {
		return nil
	}
	return privateError{
		code: c,
		err:  errors.Errorf(format, a...),
	}
}

func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	var e privateError
	if errors.As(err, &e) {
		return e.code
	}
	return codes.Unknown
}

func StackTrace(err error) string {
	var e privateError
	if errors.As(err, &e) {
		return fmt.Sprintf("%+v\n", e.err)
	}
	return ""
}
