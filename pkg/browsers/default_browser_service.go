package browsers

type DefaultBrowserService struct {
	BrowserRepository BrowserRepository
}

func (s *DefaultBrowserService) GetBrowsers() (BrowserSlice, error) {
	return s.BrowserRepository.GetBrowsers()
}

func (s *DefaultBrowserService) UpdateBrowser(b *Browser) error {
	return s.BrowserRepository.UpdateBrowser(b)
}
