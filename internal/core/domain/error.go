package domain

import "errors"

var (
	ErrInternal        = errors.New("internal error")
	ErrDataNotFound    = errors.New("no data found")
	ErrDataConflict    = errors.New("duplicate data found with unique column")
	ErrUnauthorized    = errors.New("user is unauthorized to access this resource")
	ErrForbidden       = errors.New("access to this resource is forbidden")
	ErrInvalidPassword = errors.New("invalid password for user")
)
