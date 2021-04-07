package repositories

import "errors"

var (
	ErrNoRows = errors.New("sql: no rows")
)