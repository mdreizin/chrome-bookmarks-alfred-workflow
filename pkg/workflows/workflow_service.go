package workflows

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/bookmarks"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/profiles"
)

type WorkflowService interface {
	GetBrowsers() (browsers.BrowserSlice, error)
	UpdateBrowser(*browsers.Browser) error
	GetProfiles(*browsers.Browser, string) (profiles.ProfileSlice, error)
	GetBookmarks(*browsers.Browser, string) (bookmarks.BookmarkSlice, error)
	GetWorkflowMetadata() (WorkflowMetadata, error)
}
