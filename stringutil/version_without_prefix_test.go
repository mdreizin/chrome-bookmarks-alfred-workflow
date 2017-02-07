package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionWithoutPrefix(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("1.0.0", VersionWithoutPrefix("v1.0.0"))
}
