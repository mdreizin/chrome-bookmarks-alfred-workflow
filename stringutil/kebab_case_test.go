package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKebabCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("FOO_BAR", KebabCase("foo-bar"))
}
