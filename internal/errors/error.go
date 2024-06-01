package errors

import "errors"

var (
	ErrNoPath      = errors.New("path required")
	ErrInvalidArgs = errors.New("invalid number of arguments")
)
