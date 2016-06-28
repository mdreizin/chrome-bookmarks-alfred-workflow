package stringutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDefaultIfEmpty(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Foo", DefaultIfEmpty("Foo", "Bar"))
	assert.Equal("Bar", DefaultIfEmpty("", "Bar"))
}
