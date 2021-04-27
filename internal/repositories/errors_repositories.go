package repositories

import "errors"

var (
	ErrNoResult = errors.New("sql: no rows")
	ErrOnSave   = errors.New("sql: error to save object")
	ErrOnUpdate = errors.New("sql: failed to update object")
	ErrOnDelete = errors.New("sql: failed to delete object")
)
