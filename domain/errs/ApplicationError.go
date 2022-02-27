package errs

import (
	"errors"
	"fmt"
)

type ApplicationError struct {
	OriginErr error
}

func (s ApplicationError) Error() string {
	return fmt.Errorf("application error: %w", s.OriginErr).Error()
}

func (s ApplicationError) UnWrap() error {
	return s.OriginErr
}

func IsApplicationError(err error) bool {
	var target ApplicationError
	return errors.As(err, &target)
}
