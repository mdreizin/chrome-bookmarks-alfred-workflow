package bookmarks

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonBookmarkRepository_GetBookmarks(t *testing.T) {
	test := assert.New(t)
	bookmarkRepository := JsonBookmarkRepository{Filename: "Bookmarks.json"}
	browser := browsers.Browser{Path: "test/Default"}
	bookmarkSlice, err := bookmarkRepository.GetBookmarks(&browser)

	test.NoError(err)
	test.Len(bookmarkSlice, 7)
}

func TestJsonBookmarkRepository_GetBookmarks_ReadFileError(t *testing.T) {
	test := assert.New(t)
	bookmarkRepository := JsonBookmarkRepository{Filename: ""}
	browser := browsers.Browser{Path: "test/Default"}
	bookmarkSlice, err := bookmarkRepository.GetBookmarks(&browser)

	test.Error(err)
	test.Len(bookmarkSlice, 0)
}

func TestJsonBookmarkRepository_GetBookmarks_UnmarshalError(t *testing.T) {
	test := assert.New(t)
	bookmarkRepository := JsonBookmarkRepository{Filename: "Bookmarks"}
	browser := browsers.Browser{Path: "test/Default"}
	bookmarkSlice, err := bookmarkRepository.GetBookmarks(&browser)

	test.Error(err)
	test.Len(bookmarkSlice, 0)
}
