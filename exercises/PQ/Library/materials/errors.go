package materials

import "errors"

var (
	ErrDefault       = errors.New("default error")
	ErrInvalidAccess = errors.New("invalid access")
	ErrTokenInvalid  = errors.New(" invalid token")
)
