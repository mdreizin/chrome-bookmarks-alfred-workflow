package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookmarkSlice_Match(t *testing.T) {
	const unicode = "❤"

	assert := assert.New(t)
	bookmarks := BookmarkSlice{
		Bookmark{
			Name: "github.com",
			URL:  "https://github.com",
		},
		Bookmark{
			Name: unicode,
			URL:  "https://en.wikipedia.org/wiki/Unicode",
		},
	}

	assert.Len(bookmarks.Match("github.com"), 1)
	assert.Len(bookmarks.Match("hub"), 1)
	assert.Len(bookmarks.Match(unicode), 1)
	assert.Len(bookmarks.Match("noop"), 0)
}

func TestBookmarkSlice_Sort(t *testing.T) {
	assert := assert.New(t)
	bookmarks := BookmarkSlice{
		Bookmark{
			Name: "wikipedia.org",
			URL:  "https://en.wikipedia.org/wiki/Unicode",
		},
		Bookmark{
			Name: "github.com",
			URL:  "https://github.com",
		},
	}.Sort()

	assert.Equal(Bookmark{
		Name: "github.com",
		URL:  "https://github.com",
	}, bookmarks[0])
}
