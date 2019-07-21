package gocrapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Client holds info about how to connect to clash's api
type Client struct {
	token   string
	timeout time.Duration
	baseURL string
}

const officialAPI = "https://api.royaleapi.com"

// New creates a brand new client for interacting with the api
func New(token string) (*Client, error) {
	if len(token) == 0 {
		return nil, errors.New("Token not provided")
	}
	c := Client{
		token:   token,
		timeout: 10 * time.Second,
		baseURL: officialAPI,
	}
	return &c, nil
}

func (c *Client) newClient() http.Client {
	return http.Client{Timeout: c.timeout}
}

func (c *Client) Get(item string, params url.Values) ([]byte, error) {
	request, err := http.NewRequest("GET", c.baseURL+item, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating request")
	}
	request.Header.Add(
		"Authorization",
		fmt.Sprintf("Bearer %s", c.token),
	)
	request.URL.RawQuery = params.Encode()
	hc := c.newClient()
	response, reqErr := hc.Do(request)
	if reqErr != nil {
		return nil, errors.Wrap(reqErr, "Error doing request")
	}
	// TODO: maybe this should just return a response object and handle it elsewhere
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
