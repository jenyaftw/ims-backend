package domain

import "errors"

var (
	ErrInternal                = errors.New("internal error")
	ErrDataNotFound            = errors.New("no data found")
	ErrDataConflict            = errors.New("duplicate data found with unique column")
	ErrUnauthorized            = errors.New("user is unauthorized to access this resource")
	ErrForbidden               = errors.New("access to this resource is forbidden")
	ErrUserNotFound            = errors.New("user not found")
	ErrInventoryNotFound       = errors.New("inventory not found")
	ErrInvalidPassword         = errors.New("invalid password for user")
	ErrInvalidVerificationCode = errors.New("invalid verification code")
	ErrVerificationCodeExpired = errors.New("verification code expired")
	ErrUserAlreadyVerified     = errors.New("user is already verified")

	ErrMissingAuthHeader = errors.New("missing `Authorization` header")
	ErrInvalidAuthToken  = errors.New("invalid token in `Authorization` header")
	ErrInvalidTokenType  = errors.New("invalid token type")
	ErrUserNotVerified   = errors.New("user is not verified")
)
