package autoria

import "errors"

var (
	// ErrInvalidAPIKey is returned when API key is invalid
	ErrInvalidAPIKey = errors.New("invalid api key")

	// ErrTooManyRequests is returned when too many requests
	ErrTooManyRequests = errors.New("too many requests")

	// ErrInvalidRequest is returned when request is invalid
	ErrInvalidRequest = errors.New("invalid request")
)
