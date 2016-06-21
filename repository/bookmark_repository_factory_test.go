package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewBookmarkRepository(t *testing.T) {
	assert := assert.New(t)
	bookmarkRepository := NewBookmarkRepository("Bookmarks.json")

	assert.IsType(&JsonBookmarkRepository{}, bookmarkRepository)
}
