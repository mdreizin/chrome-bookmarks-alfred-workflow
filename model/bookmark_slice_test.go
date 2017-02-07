package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookmarkSlice_Match(t *testing.T) {
	const unicode = "❤"

	test := assert.New(t)
	bookmarks := BookmarkSlice{
		Bookmark{
			Name: "github.com",
			URL:  "https://github.com",
		},
		Bookmark{
			Name: unicode,
			URL:  "https://en.wikipedia.org/wiki/Unicode",
		},
		Bookmark{
			Name: "react",
			URL:  "https://facebook.github.io/react/",
		},
		Bookmark{
			Name: "react redux",
			URL:  "http://redux.js.org/docs/basics/UsageWithReact.html",
		},
	}

	test.Len(bookmarks.Match("github.com"), 1)
	test.Len(bookmarks.Match("github"), 2)
	test.Len(bookmarks.Match(unicode), 1)
	test.Len(bookmarks.Match("noop"), 0)
	test.Len(bookmarks.Match("react"), 2)
	test.Len(bookmarks.Match("react redux"), 1)
	test.Len(bookmarks.Match("React Redux"), 1)
}

func TestBookmarkSlice_Sort(t *testing.T) {
	test := assert.New(t)
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

	test.Equal(Bookmark{
		Name: "github.com",
		URL:  "https://github.com",
	}, bookmarks[0])
}
