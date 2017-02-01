package provider

import (
	"strings"
)

const (
	helloFresh = "hellofresh"
)

func Create(providerURL string, name string) Provider {
	switch strings.ToLower(name) {
	case helloFresh:
		return NewHelloFresh(providerURL)
	default:
		return nil
	}
}
