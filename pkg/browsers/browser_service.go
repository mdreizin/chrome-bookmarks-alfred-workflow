package browsers

type BrowserService interface {
	GetBrowsers() (BrowserSlice, error)
	UpdateBrowser(*Browser) error
}
