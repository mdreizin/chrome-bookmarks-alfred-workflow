package repository

import (
	"encoding/json"
	"github.com/mdreizin/chrome-bookmarks-alfred-workflow/model"
	"io/ioutil"
)

type bookmarkTree struct {
	model.Bookmark
	Type     string         `json:"type"`
	Children []bookmarkTree `json:"children,omitempty"`
}

type bookmarkAux struct {
	Roots struct {
		BookmarkBar bookmarkTree `json:"bookmark_bar"`
		Other       bookmarkTree `json:"other"`
		Synced      bookmarkTree `json:"synced"`
	} `json:"roots"`
}

func (t bookmarkTree) Walk(f func(bookmarkTree)) {
	for _, v := range t.Children {
		f(v)

		if len(v.Children) > 0 {
			v.Walk(f)
		}
	}
}

type JsonBookmarkRepository struct {
	filename string
}

func (r *JsonBookmarkRepository) GetBookmarks(browser model.Browser) (model.BookmarkSlice, error) {
	bookmarks := model.BookmarkSlice{}
	filename := browser.FullPathFor(browser.ProfileName, r.filename)

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return bookmarks, err
	}

	aux := bookmarkAux{}
	err = json.Unmarshal(bytes, &aux)

	if err != nil {
		return bookmarks, err
	}

	walk := func(t bookmarkTree) {
		if t.Type == "url" {
			bookmarks = bookmarks.Add(model.Bookmark{
				Name:    t.Name,
				URL:     t.URL,
				IconURL: browser.IconURL,
			})
		}
	}

	aux.Roots.BookmarkBar.Walk(walk)
	aux.Roots.Synced.Walk(walk)
	aux.Roots.Other.Walk(walk)

	return bookmarks, err
}
