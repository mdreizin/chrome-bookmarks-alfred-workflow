package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmarkService(t *testing.T) {
	assert := assert.New(t)
	bookmarkService := NewBookmarkService(nil)

	assert.IsType(&DefaultBookmarkService{}, bookmarkService)
}
