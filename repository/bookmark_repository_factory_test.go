package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBookmarkRepository(t *testing.T) {
	test := assert.New(t)
	bookmarkRepository := NewBookmarkRepository("Bookmarks.json")

	test.IsType(&JsonBookmarkRepository{}, bookmarkRepository)
}
