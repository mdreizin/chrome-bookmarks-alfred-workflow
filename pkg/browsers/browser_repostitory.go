package browsers

type BrowserRepository interface {
	GetBrowsers() (BrowserSlice, error)
	UpdateBrowser(*Browser) error
}
