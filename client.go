package gocrapi

import (
	"encoding/json"
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

func (c *Client) get(item string, params url.Values) (*http.Response, error) {
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
	return response, nil
}

func handleResponse(response *http.Response, result interface{}) error {
	defer response.Body.Close()
	readbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "Error while reading body from response")
	}
	if response.StatusCode >= 400 {
		return fmt.Errorf("Error from the api [%s]: %s", response.Status, readbody)
	}
	err = json.Unmarshal(readbody, result)
	if err != nil {
		return errors.Wrap(err, "Error while parsing the response into json")
	}
	return nil

}
