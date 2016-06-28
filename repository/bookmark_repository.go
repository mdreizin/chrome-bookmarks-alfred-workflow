package repository

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"

type BookmarkRepository interface {
	GetBookmarks(model.Browser) (model.BookmarkSlice, error)
}
