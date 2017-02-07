package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBookmarkService(t *testing.T) {
	test := assert.New(t)
	bookmarkService := NewBookmarkService(nil)

	test.IsType(&DefaultBookmarkService{}, bookmarkService)
}
