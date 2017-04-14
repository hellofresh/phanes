package provider

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"encoding/json"

	log "github.com/Sirupsen/logrus"
	"github.com/hellofresh/phanes/pkg/generator"
	uuid "github.com/satori/go.uuid"
)

// HelloFreshClient represents the hellofresh client
type HelloFreshClient struct {
	ID          uuid.UUID `json:"id"`
	Secret      string    `json:"secret"`
	Extra       string    `json:"extra"`
	RedirectURI string    `json:"redirect_uri"`
}

// GetID retrieves the client's ID
func (c *HelloFreshClient) GetID() string {
	return c.ID.String()
}

// GetSecret retreives the client's secret
func (c *HelloFreshClient) GetSecret() string {
	return c.Secret
}

// HelloFresh represents a provider
type HelloFresh struct {
	url    string
	client *http.Client
}

// NewHelloFresh creates a new instance of HelloFresh
func NewHelloFresh(url string) *HelloFresh {
	return &HelloFresh{
		url,
		&http.Client{
			Timeout: time.Second * 10,
		}}
}

// Create a client
func (p *HelloFresh) Create(name string, redirectURI string) (Client, error) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	client := new(HelloFreshClient)
	client.ID = uuid.NewV4()
	client.Secret = fmt.Sprintf("%x", generator.GenerateSecret(client.ID.String()))
	client.Extra = name
	client.RedirectURI = redirectURI
	s := client.GetSecret()
	log.Debug(s)

	jsonStr, err := json.Marshal(&client)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", p.url, bytes.NewBuffer(jsonStr))
	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		log.Debug(string(bodyBytes))
		return nil, errors.New("Client not created")
	}

	return client, nil
}

// Remove a client
func (p *HelloFresh) Remove(id string) error {
	req, err := http.NewRequest("DELETE", p.url+id, nil)
	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return errors.New("Client not found")
	case http.StatusNoContent:
		return nil
	default:
		return errors.New("Client not deleted")
	}
}