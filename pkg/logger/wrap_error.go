package logger

import "github.com/pkg/errors"

func WrapError(err error) error {
	return errors.Wrap(err, err.Error())
}
