package services

import "errors"

var (
	ErrCategoryNotFound = errors.New("service: category not found")
	ErrCategoryUpdate   = errors.New("service: failed to update category")
)
