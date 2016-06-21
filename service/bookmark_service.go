package service

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"

type BookmarkService interface {
	GetBrowsers() (model.BrowserSlice, error)
	GetProfiles(model.Browser, string) (model.ProfileSlice, error)
	GetBookmarks(model.Browser, string) (model.BookmarkSlice, error)
}
