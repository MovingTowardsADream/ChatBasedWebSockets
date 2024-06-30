package repository_error

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")

	ErrNotFound = errors.New("not found")
)
