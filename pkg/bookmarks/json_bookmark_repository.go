package bookmarks

import (
	"encoding/json"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/pkg/browsers"
	"io/ioutil"
)

type bookmarkTree struct {
	Bookmark
	Type     string         `json:"type"`
	Name     string         `json:"name"`
	Children []bookmarkTree `json:"children,omitempty"`
}

type bookmarkAux struct {
	Roots struct {
		BookmarkBar bookmarkTree `json:"bookmark_bar"`
		Other       bookmarkTree `json:"other"`
		Synced      bookmarkTree `json:"synced"`
	} `json:"roots"`
}

func (t bookmarkTree) Walk(f func(*bookmarkTree)) {
	if t.Type == "folder" && t.Name != "" {
		if t.Path == nil {
			t.Path = []string{}
		}

		t.Path = append(t.Path, t.Name)
	}

	for _, child := range t.Children {
		child.Path = t.Path

		f(&child)

		if len(child.Children) > 0 {
			child.Walk(f)
		}
	}
}

type JsonBookmarkRepository struct {
	Filename string
}

func (r *JsonBookmarkRepository) GetBookmarks(browser *browsers.Browser) (BookmarkSlice, error) {
	filename := browser.ResolvePath(browser.ProfileName, r.Filename)
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	aux := bookmarkAux{}
	err = json.Unmarshal(bytes, &aux)

	if err != nil {
		return nil, err
	}

	bookmarkSlice := BookmarkSlice{}

	walk := func(t *bookmarkTree) {
		if t.Type == "url" {
			bookmarkSlice = bookmarkSlice.Add(Bookmark{
				Name:    t.Name,
				URL:     t.URL,
				Path:    t.Path,
				IconURL: browser.IconURL,
			})
		}
	}

	aux.Roots.BookmarkBar.Walk(walk)
	aux.Roots.Synced.Walk(walk)
	aux.Roots.Other.Walk(walk)

	return bookmarkSlice, nil
}
