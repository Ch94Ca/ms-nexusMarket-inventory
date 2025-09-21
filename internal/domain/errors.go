package domain

import "errors"

var (
	ErrInvalidCategoryName = errors.New("category name cannot be empty")
	ErrCategoryNotFound    = errors.New("category not found")
)
