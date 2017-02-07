package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultIfEmpty(t *testing.T) {
	test := assert.New(t)

	test.Equal("Foo", DefaultIfEmpty("Foo", "Bar"))
	test.Equal("Bar", DefaultIfEmpty("", "Bar"))
}
