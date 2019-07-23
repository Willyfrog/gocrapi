package gocrapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const clanTag = "%23Q2G2UOG"

func TestGetClan(t *testing.T) {
	c, clerr := New(TOKEN)

	assert.Nil(t, clerr)

	warclan, err := c.GetClan(clanTag)

	assert.NotNil(t, warclan)
	assert.Nil(t, err)
}
