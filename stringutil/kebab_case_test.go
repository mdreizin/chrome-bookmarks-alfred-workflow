package stringutil

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestKebabCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("FOO_BAR", KebabCase("foo-bar"))
}
