package service

import "github.com/pkg/errors"

var (
	ErrNotAuthorized = errors.New("not authorized")
)
