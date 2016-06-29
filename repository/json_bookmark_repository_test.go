package repository

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonBookmarkRepository_GetBookmarks(t *testing.T) {
	assert := assert.New(t)
	repository := JsonBookmarkRepository{filename: "Bookmarks.json"}
	browser := model.Browser{Path: "testdata/Default"}
	bookmarks, err := repository.GetBookmarks(browser)

	assert.NoError(err)
	assert.Len(bookmarks, 7)
}

func TestJsonBookmarkRepository_GetBookmarks_ReadFileError(t *testing.T) {
	assert := assert.New(t)
	repository := JsonBookmarkRepository{filename: ""}
	browser := model.Browser{Path: "testdata/Default"}
	bookmarks, err := repository.GetBookmarks(browser)

	assert.Error(err)
	assert.Len(bookmarks, 0)
}

func TestJsonBookmarkRepository_GetBookmarks_UnmarshalError(t *testing.T) {
	assert := assert.New(t)
	repository := JsonBookmarkRepository{filename: "Bookmarks"}
	browser := model.Browser{Path: "testdata/Default"}
	bookmarks, err := repository.GetBookmarks(browser)

	assert.Error(err)
	assert.Len(bookmarks, 0)
}
