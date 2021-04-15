package repositories

import "errors"

var (
	ErrNoResult = errors.New("sql: no rows")
)
