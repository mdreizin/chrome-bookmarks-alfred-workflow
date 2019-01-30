package bookmarks

import "github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"

type DefaultBookmarkService struct {
	BookmarkRepository BookmarkRepository
}

func (s *DefaultBookmarkService) GetBookmarks(browser *browsers.Browser) (BookmarkSlice, error) {
	return s.BookmarkRepository.GetBookmarks(browser)
}
