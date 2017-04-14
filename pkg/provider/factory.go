package provider

import (
	"strings"
)

const (
	helloFresh = "hellofresh"
)

// Create the correct provider for a given name
func Create(name string, providerURL string) (Provider, error) {
	switch strings.ToLower(name) {
	case helloFresh:
		return NewHelloFresh(providerURL), nil
	default:
		return nil, ErrInvalidProvider
	}
}
