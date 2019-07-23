package gocrapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const clanTag = "#Q2G2U0G"

func TestGetClan(t *testing.T) {
	c, clerr := New(TOKEN)

	assert.Nil(t, clerr)

	warclan, err := c.GetClan(NewTag(clanTag))

	assert.NotNil(t, warclan)
	assert.Nil(t, err)
}
