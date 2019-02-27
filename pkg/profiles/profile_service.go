package profiles

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"

type ProfileService interface {
	GetProfiles(*browsers.Browser) (ProfileSlice, error)
}
