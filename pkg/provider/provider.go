package provider

// Client holds the basic methods for a client
type Client interface {
	GetID() string
	GetSecret() string
}

// Provider holds the basic methods for a provider
type Provider interface {
	Create(name string, redirectURI string) (Client, error)
	Remove(id string) error
}
