package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionWithoutPrefix(t *testing.T) {
	test := assert.New(t)

	test.Equal("1.0.0", VersionWithoutPrefix("v1.0.0"))
}
