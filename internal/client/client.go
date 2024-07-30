package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// HostURL - Default Administration URL.
const AuthServerURL string = "https://auth.quortex.io"
const HostURL string = "https://api.quortex.io"

type Client struct {
	AuthServerURL string
	HostURL       string
	HTTPClient    *http.Client
	Token         string
	Auth          AuthStruct
}

type AuthStruct struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewClient(auth_server, host, client_id, client_secret *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Administration URL
		HostURL:       HostURL,
		AuthServerURL: AuthServerURL,
		Auth: AuthStruct{
			ClientId:     *client_id,
			ClientSecret: *client_secret,
			GrantType:    "client_credentials",
		},
	}

	if auth_server != nil {
		c.AuthServerURL = *auth_server
	}

	if host != nil {
		c.HostURL = *host
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = "Bearer " + ar.AccessToken

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.Token != "" {
		req.Header.Set("Authorization", c.Token)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(res)
	log.Println(res.StatusCode)
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
