package repository

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonBookmarkRepository_GetBookmarks(t *testing.T) {
	test := assert.New(t)
	repository := JsonBookmarkRepository{filename: "Bookmarks.json"}
	browser := model.Browser{Path: "testdata/Default"}
	bookmarks, err := repository.GetBookmarks(browser)

	test.NoError(err)
	test.Len(bookmarks, 7)
}

func TestJsonBookmarkRepository_GetBookmarks_ReadFileError(t *testing.T) {
	test := assert.New(t)
	repository := JsonBookmarkRepository{filename: ""}
	browser := model.Browser{Path: "testdata/Default"}
	bookmarks, err := repository.GetBookmarks(browser)

	test.Error(err)
	test.Len(bookmarks, 0)
}

func TestJsonBookmarkRepository_GetBookmarks_UnmarshalError(t *testing.T) {
	test := assert.New(t)
	repository := JsonBookmarkRepository{filename: "Bookmarks"}
	browser := model.Browser{Path: "testdata/Default"}
	bookmarks, err := repository.GetBookmarks(browser)

	test.Error(err)
	test.Len(bookmarks, 0)
}
