package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBookmarkService(t *testing.T) {
	assert := assert.New(t)
	bookmarkService := NewBookmarkService(nil)

	assert.IsType(&DefaultBookmarkService{}, bookmarkService)
}
