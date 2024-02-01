package domain

import "errors"

var (
	ErrInternal     = errors.New("internal error")
	ErrDataNotFound = errors.New("no data with such input found")
	ErrDataConflict = errors.New("duplicate data found with unique column")
	ErrUnauthorized = errors.New("user is unauthorized to access this resource")
	ErrForbidden    = errors.New("access to thisresource is forbidden")
)
