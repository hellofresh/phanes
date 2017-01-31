package provider

type Client interface {
	GetID() string
	GetSecret() string
}

type Provider interface {
	Create(name string, redirectURI string) (Client, error)
	Delete(id string) error
}
