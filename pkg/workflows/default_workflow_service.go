package workflows

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/bookmarks"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/profiles"
)

type DefaultWorkflowService struct {
	BrowserService     browsers.BrowserService
	ProfileService     profiles.ProfileService
	BookmarkService    bookmarks.BookmarkService
	WorkflowRepository WorkflowRepository
}

func (s *DefaultWorkflowService) GetBrowsers() (browsers.BrowserSlice, error) {
	return s.BrowserService.GetBrowsers()
}

func (s *DefaultWorkflowService) UpdateBrowser(b *browsers.Browser) error {
	return s.BrowserService.UpdateBrowser(b)
}

func (s *DefaultWorkflowService) GetProfiles(b *browsers.Browser, query string) (profiles.ProfileSlice, error) {
	profileSlice, err := s.ProfileService.GetProfiles(b)

	if err != nil {
		return nil, err
	}

	if query != "" {
		profileSlice = profileSlice.Match(query)
	}

	return profileSlice.Sort(), nil
}

func (s *DefaultWorkflowService) GetBookmarks(browser *browsers.Browser, query string) (bookmarks.BookmarkSlice, error) {
	profileSlice, err := s.ProfileService.GetProfiles(browser)

	if err != nil {
		return nil, err
	}

	profileName := profiles.DefaultProfileName

	if browser.ProfileName == "" {
		if profile, err := profileSlice.FirstActive(); err == nil && profile.Name != "" {
			profileName = profile.Name
		}
	}

	browser.ProfileName = profileName

	bookmarkSlice, err := s.BookmarkService.GetBookmarks(browser)

	if err != nil {
		return nil, err
	}

	if query != "" {
		bookmarkSlice = bookmarkSlice.Match(query)
	}

	return bookmarkSlice.Sort(), nil
}

func (s *DefaultWorkflowService) GetWorkflowMetadata() (WorkflowMetadata, error) {
	return s.WorkflowRepository.GetWorkflowMetadata()
}
