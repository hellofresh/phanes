package provider

import "errors"

var (
	// ErrInvalidProvider is used when the provider is invalid
	ErrInvalidProvider = errors.New("invalid provider selected, please choose a valid provider")
)
