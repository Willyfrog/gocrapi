package gocrapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const clanTag = "#Q2G2U0G"

func setupResponseServer(t *testing.T, requestedURL string, status int, body string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, req *http.Request) {
		// Test request parameters
		require.Equal(t, req.URL.String(), requestedURL)
		response.WriteHeader(status)
		response.Header().Set("Content-Type", "application/json")
		// Send response to be tested
		io.WriteString(response, body)
	}))
	return server
}

func NewTestClient(srv *httptest.Server) *Client {
	c, _ := New(TOKEN)
	c.client = srv.Client()
	return c
}
// TODO: add test with 404

// TODO: add test with success

// TODO: add test with unathorized

// TODO: remove this test
func TestGetClan(t *testing.T) {

	t.Run("accessDenied", func(t *testing.T) {
		response := "{\"reason\":\"accessDenied\",\"message\":\"Invalid authorization\"}"
		server := setupResponseServer(t, "/clans/%23Q2G2U0G/currentwar", http.StatusForbidden, response)
		// Close the server when test finishes
		defer server.Close()
		c := NewTestClient(server)
		c.baseURL = "http://localhost"
		clan, err := c.GetClan(NewTag(clanTag))
		assert.Nil(t, clan)
		assert.NotNil(t, err)
		errCause := errors.Cause(err)
		require.NotNil(t, errCause)
		require.Equal(t, errCause.(type), *url.Error)
		originalErr := errCause
		assert.Nil(t, originalErr)
	})

	t.Run("Running agains real server, so it might fail", func(t *testing.T) {
		c, clerr := New(TOKEN)

		assert.Nil(t, clerr)

		warclan, err := c.GetClan(NewTag(clanTag))

		assert.NotNil(t, warclan)
		assert.Nil(t, err)

		//assert.Nil(t, warclan) // this should break
	})
}
