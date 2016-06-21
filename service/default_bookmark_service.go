package service

import (
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/repository"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
)

type DefaultBookmarkService struct {
	config					map[string]string
	browserRepository 		repository.BrowserRepository
	profileRepository		repository.ProfileRepository
	bookmarkRepository 		repository.BookmarkRepository
}

func (s *DefaultBookmarkService) BrowserRepository() repository.BrowserRepository {
	if s.browserRepository == nil {
		s.browserRepository = repository.NewBrowserRepository(s.config["browser-file"])
	}

	return s.browserRepository
}

func (s *DefaultBookmarkService) ProfileRepository() repository.ProfileRepository {
	if s.profileRepository == nil {
		s.profileRepository = repository.NewProfileRepository(s.config["profile-file"])
	}

	return s.profileRepository
}

func (s *DefaultBookmarkService) BookmarkRepository() repository.BookmarkRepository {
	if s.bookmarkRepository == nil {
		s.bookmarkRepository = repository.NewBookmarkRepository(s.config["bookmark-file"])
	}

	return s.bookmarkRepository
}

func (s *DefaultBookmarkService) GetBrowsers() (model.BrowserSlice, error) {
	return s.BrowserRepository().GetBrowsers()
}

func (s *DefaultBookmarkService) GetProfiles(b model.Browser, query string) (model.ProfileSlice, error) {
	profiles, err := s.ProfileRepository().GetProfiles(b)

	if err == nil {
		if (query != "") {
			profiles = profiles.Match(query)
		}

		profiles = profiles.Sort()
	}

	return profiles, err
}

func (s *DefaultBookmarkService) GetBookmarks(browser model.Browser, query string) (model.BookmarkSlice, error) {
	bookmarks := model.BookmarkSlice{}
	profiles, _ := s.GetProfiles(browser, "")

	if browser.ProfileName == "" {
		profile, _ := profiles.FirstActive()

		if profile.Name != "" {
			browser.ProfileName = profile.Name
		}
	}

	bookmarks, err := s.BookmarkRepository().GetBookmarks(browser)

	if err == nil {
		if (query != "") {
			bookmarks = bookmarks.Match(query)
		}

		bookmarks = bookmarks.Sort()
	}

	return bookmarks, err
}
