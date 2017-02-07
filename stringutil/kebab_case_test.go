package stringutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKebabCase(t *testing.T) {
	test := assert.New(t)

	test.Equal("FOO_BAR", KebabCase("foo-bar"))
}
