package v1

import (
	"errors"
)

var (
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrCannotParseToken  = errors.New("cannot parse token")

	ErrIdNotFound    = errors.New("user id not found")
	ErrIdInvalidType = errors.New("user id is of invalid type")
)
