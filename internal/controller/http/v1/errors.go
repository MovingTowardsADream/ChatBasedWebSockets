package v1

import (
	"errors"
)

var (
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrCannotParseToken  = errors.New("cannot parse token")
)
