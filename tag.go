package gocrapi

import (
	"fmt"
	"strings"
)

// Tag contains either a clan identifier or a user identifier
type Tag string

const urlHash = "%23"

// NewTag generates a new tag from a string by trimming any possible prefix and storing the raw value
func NewTag(id string) *Tag {
	id = strings.TrimPrefix(id, urlHash)
	t := Tag(strings.TrimPrefix(id, "#"))
	return &t
}

func (t *Tag) String() string {
	return fmt.Sprintf("#%s", string(*t))
}

// URLEncode convert tag into something a browser can understand
func (t *Tag) URLEncode() string {
	return fmt.Sprintf("%s%s", urlHash, string(*t))
}
