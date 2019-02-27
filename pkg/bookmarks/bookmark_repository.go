package bookmarks

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"

type BookmarkRepository interface {
	GetBookmarks(*browsers.Browser) (BookmarkSlice, error)
}
