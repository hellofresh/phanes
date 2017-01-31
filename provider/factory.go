package provider

import (
	"strings"
)

const (
	helloFresh = "hellofresh"
)

func Create(name string) Provider {
	switch strings.ToLower(name) {
	case helloFresh:
		return NewHelloFresh()
	default:
		return nil
	}
}
