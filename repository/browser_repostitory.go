package repository

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"

type BrowserRepository interface {
	GetBrowsers() (model.BrowserSlice, error)
	UpdateBrowser(model.Browser) error
}
