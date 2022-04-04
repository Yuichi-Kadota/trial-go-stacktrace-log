package repository

import (
	"trial-go-stacktrace/codes"
	"trial-go-stacktrace/internal/errors"
)

type UserImpl struct{}

func (r UserImpl) GetUser() error {
	return errors.Errorf(codes.Database, "failed to get user")
}
