package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBookmarkRepository(t *testing.T) {
	assert := assert.New(t)
	bookmarkRepository := NewBookmarkRepository("Bookmarks.json")

	assert.IsType(&JsonBookmarkRepository{}, bookmarkRepository)
}
