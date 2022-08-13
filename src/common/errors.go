package common

import "errors"

var (
	ErrNotFound       = errors.New("Resource Not Found")
	ErrBadRequest     = errors.New("Bad Request")
	ErrInternal       = errors.New("Internal Error")
	ErrCreationFailed = errors.New("Failed to create the resource")
)
