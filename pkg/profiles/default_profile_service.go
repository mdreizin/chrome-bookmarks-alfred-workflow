package profiles

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"

type DefaultProfileService struct {
	ProfileRepository ProfileRepository
}

func (s *DefaultProfileService) GetProfiles(browser *browsers.Browser) (ProfileSlice, error) {
	return s.ProfileRepository.GetProfiles(browser)
}
