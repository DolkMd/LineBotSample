package errs

import (
	"errors"
	"fmt"
)

type streamerErrorType string

const (
	StreamerErrorUnKnown  streamerErrorType = ""
	StreamerErrorExternal streamerErrorType = "external"
	StreamerErrorRequest  streamerErrorType = "request"
)

func StreamerErrorType(err error) streamerErrorType {
	var target StreamerError
	if errors.As(err, &target) {
		return target.ErrType
	}
	return StreamerErrorUnKnown
}

type StreamerError struct {
	OriginErr error
	ErrType   streamerErrorType
}

func IsStreamerError(err error) bool {
	var target StreamerError
	return errors.As(err, &target)
}

func (s StreamerError) Error() string {
	return fmt.Errorf("streamer error: %w", s.OriginErr).Error()
}

func (s StreamerError) UnWrap() error {
	return s.OriginErr
}
