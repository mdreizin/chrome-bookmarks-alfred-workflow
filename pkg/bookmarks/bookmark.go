package bookmarks

type Bookmark struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string
	Path    []string
}
