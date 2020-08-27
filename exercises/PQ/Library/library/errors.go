package library

import "errors"

var (
	ErrDefault       = errors.New("default error")
	ErrInvalidAccess = errors.New("invalid access")
)
