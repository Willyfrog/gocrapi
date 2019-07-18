package gocrapi

import (
	"net/http"
	"time"
	"errors"
)

// Client holds info about how to connect to clash's api
type Client struct {
	token string
	timeout time.Duration
	baseUrl string
}

const officialAPI = "https://api.royaleapi.com"

// New creates a brand new client for interacting with the api
func New(token string) (*Client, error) {
	if len(token) == 0 {
		return nil, errors.New("Token not provided")
	}
	c := Client{
		token: token,
		timeout: 10 * time.Second,
		baseUrl: officialAPI,
	}
	return &c, nil
}

func (Client *c) newClient() http.Client {
	return http.Client{Timeout: c.timeout}
}

func (Client *c) Get(item string, params url.Values) ([]bytes, error) (
	request, err := http.NewRequest("GET", c.baseUrl+item, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add(
		"Authorization", 
		fmt.Sprintf("Bearer %s", c.token)
	)
	request.URL.RawQuery = params.Encode
	hc := c.newClient()
	response, reqErr := hc.Do(request)
	if reqErr != nil {
		return nil, err
	}
	defer response.Body.Close()
	return response.Body, nil
)
