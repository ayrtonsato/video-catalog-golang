package services

import "errors"

var (
	ErrNotFound     = errors.New("service: object not found")
	ErrUpdateFailed = errors.New("service: failed to update object")
	ErrSaveFailed   = errors.New("service: failed to save object")
)
