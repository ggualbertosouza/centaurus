package centaurus

import "errors"

var (
	ErrInvalidAlgorithm       = errors.New("invalid algorithm")
	ErrInvalidConcurrencyMode = errors.New("invalid concurrency mode")
	ErrMinRate                = errors.New("rate must be greater than zero")
	ErrInvalidBurst           = errors.New("burst must be greater than or equal to rate")
)
