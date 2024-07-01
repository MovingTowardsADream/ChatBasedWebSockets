package usecase

import (
	"context"
	"errors"
)

var (
	ErrTimeout = context.DeadlineExceeded

	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)
