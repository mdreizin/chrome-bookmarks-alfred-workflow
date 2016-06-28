package repository

func NewBrowserRepository(filename string) BrowserRepository {
	return &YmlBrowserRepository{filename: filename}
}
