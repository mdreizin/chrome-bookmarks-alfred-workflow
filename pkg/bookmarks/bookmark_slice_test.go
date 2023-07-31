package bookmarks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookmarkSlice_Match(t *testing.T) {
	const heartSymbol = "❤"
	const bugSymbol = "🐛"

	test := assert.New(t)
	bookmarkSlice := BookmarkSlice{
		&Bookmark{
			Name: "Golang",
			URL:  "https://golang.org",
		},
		&Bookmark{
			Name: heartSymbol,
			URL:  "https://unicode-table.com/en/1F5A4/",
		},
		&Bookmark{
			Name: bugSymbol,
			URL:  "https://unicode-table.com/en/1F41B/",
		},
		&Bookmark{
			Name: "Rust",
			URL:  "https://www.rust-lang.org",
		},
	}

	test.Len(bookmarkSlice.Match("Go"), 1)
	test.Len(bookmarkSlice.Match("Rust"), 1)
	test.Len(bookmarkSlice.Match(heartSymbol), 1)
	test.Len(bookmarkSlice.Match(bugSymbol), 1)
	test.Len(bookmarkSlice.Match("Java"), 0)
	test.Len(bookmarkSlice.Match("golang.org"), 1)
}

func TestBookmarkSlice_Sort(t *testing.T) {
	test := assert.New(t)
	bookmarkSlice := BookmarkSlice{
		&Bookmark{
			Name: "Rust",
			URL:  "https://www.rust-lang.org",
		},
		&Bookmark{
			Name: "Golang",
			URL:  "https://golang.org",
		},
	}.Sort()

	test.Equal(&Bookmark{
		Name: "Golang",
		URL:  "https://golang.org",
	}, bookmarkSlice[0])
}
