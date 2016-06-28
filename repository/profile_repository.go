package repository

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"

type ProfileRepository interface {
	GetProfiles(model.Browser) (model.ProfileSlice, error)
}
